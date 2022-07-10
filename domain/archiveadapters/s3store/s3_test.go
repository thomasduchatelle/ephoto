package s3store

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/thomasduchatelle/dphoto/domain/archive"
	"io"
	"strings"
	"testing"
	"time"
)

var awsSession = session.Must(session.NewSession(&aws.Config{
	Credentials:      credentials.NewStaticCredentials("localstack", "localstack", ""),
	Endpoint:         aws.String("http://localhost:4566"),
	Region:           aws.String("eu-west-1"),
	S3ForcePathStyle: aws.Bool(true),
}))

func createSess(purpose string) *store {
	storeInterface := Must(New(awsSession, fmt.Sprintf("dphoto-unit-archive-%s-%s", purpose, time.Now().Format("20060102150405"))))

	adapter := storeInterface.(*store)

	_, err := adapter.s3.CreateBucket(&s3.CreateBucketInput{Bucket: &adapter.bucketName})
	if err != nil {
		panic(err)
	}

	return adapter
}

func TestUpload(t *testing.T) {
	adapter := createSess("upload")

	for _, name := range []string{"/unittest/2021/img-2021-1.jpg", "/unittest/2021/img-2021-1_01.jpg", "/unittest/2021/img-002.jpg"} {
		_, err := adapter.s3.PutObject(&s3.PutObjectInput{
			Body:   aws.ReadSeekCloser(strings.NewReader("content of " + name)),
			Bucket: &adapter.bucketName,
			Key:    aws.String(name),
		})
		if err != nil {
			panic(err)
		}
	}

	type args struct {
		key     archive.DestructuredKey
		content io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"it should upload a file, without special case", args{newKey("unittest/2020/img-2020-1", ".jpg"), strings.NewReader("foo")}, "unittest/2020/img-2020-1.jpg"},
		{"it should find a different key to avoid a clash", args{newKey("unittest/2021/img-002", ".jpg"), strings.NewReader("foo")}, "unittest/2021/img-002_01.jpg"},
		{"it should use a counter to find a different key to not override existing files", args{newKey("unittest/2021/img-2021-1", ".jpg"), strings.NewReader("foo")}, "unittest/2021/img-2021-1_02.jpg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := adapter.Upload(tt.args.key, tt.args.content)
			if a.NoError(err, tt.name) {
				a.Equal(tt.want, got, tt.name)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	adapter := createSess("copy")

	for _, name := range []string{"unittest/2021/img-2021-1.jpg", "unittest/2021/img-2021-2.jpg"} {
		_, err := adapter.s3.PutObject(&s3.PutObjectInput{
			Body:   aws.ReadSeekCloser(strings.NewReader("content of " + name)),
			Bucket: &adapter.bucketName,
			Key:    aws.String(name),
		})
		if err != nil {
			panic(err)
		}
	}

	tests := []struct {
		name    string
		origin  string
		dest    archive.DestructuredKey
		want    string
		wantErr bool
	}{
		{
			name:   "it should copy a file to requested name",
			origin: "/unittest/2021/img-2021-1.jpg",
			dest: archive.DestructuredKey{
				Prefix: "unittest/2021-q1/img-2021-1",
				Suffix: ".jpg",
			},
			want: "unittest/2021-q1/img-2021-1.jpg",
		},
		{
			name:   "it should move a file to an available name",
			origin: "/unittest/2021/img-2021-1.jpg",
			dest: archive.DestructuredKey{
				Prefix: "unittest/2021/img-2021-2",
				Suffix: ".jpg",
			},
			want: "unittest/2021/img-2021-2_01.jpg",
		},
		{
			name:   "it should return an error if the source file doesn't exist",
			origin: "foobar",
			dest: archive.DestructuredKey{
				Prefix: "unittest/2021/img-2021-1",
				Suffix: ".jpg",
			},
			wantErr: true,
		},
		{
			name:   "it should not allow prefix starting wity '/'",
			origin: "foobar",
			dest: archive.DestructuredKey{
				Prefix: "/foobar",
				Suffix: ".jpg",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := adapter.Copy(tt.origin, tt.dest)

			if !tt.wantErr && assert.NoError(t, err, tt.name) {
				assert.Equal(t, tt.want, got, tt.name)

				_, err = adapter.s3.GetObject(&s3.GetObjectInput{
					Bucket: &adapter.bucketName,
					Key:    aws.String(got),
				})
				assert.NoError(t, err, tt.name)

			} else if tt.wantErr {
				assert.Error(t, err, tt.name)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	adapter := createSess("delete")

	for _, name := range []string{"unittest/2021/img-2021-1.jpg", "unittest/2021/img-2021-2.jpg", "unittest/2021/img-2021-3.jpg", "unittest/2021/img-2021-4.jpg"} {
		_, err := adapter.s3.PutObject(&s3.PutObjectInput{
			Body:   aws.ReadSeekCloser(strings.NewReader("content of " + name)),
			Bucket: &adapter.bucketName,
			Key:    aws.String(name),
		})
		if err != nil {
			panic(err)
		}
	}

	tests := []struct {
		name string
		ids  []string
	}{
		{"it should delete one file", []string{"unittest/2021/img-2021-1.jpg"}},
		{"it should delete several files", []string{"unittest/2021/img-2021-2.jpg", "unittest/2021/img-2021-3.jpg"}},
		{"it should not fail when file is already gone", []string{"foobar"}},
		{"it should not fail when deleting several files including one that didn't exist", []string{"unittest/2021/img-2021-4.jpg", "foobar"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := adapter.Delete(tt.ids)

			if assert.NoError(t, err, tt.name) {
				for _, key := range tt.ids {
					_, err = adapter.s3.GetObject(&s3.GetObjectInput{
						Bucket: &adapter.bucketName,
						Key:    &key,
					})

					aerr, ok := err.(awserr.Error)
					assert.True(t, ok, tt.name, key, err)
					assert.Equal(t, s3.ErrCodeNoSuchKey, aerr.Code(), tt.name, key)
				}
			}
		})
	}
}

func newKey(prefix, suffix string) archive.DestructuredKey {
	return archive.DestructuredKey{
		Prefix: prefix,
		Suffix: suffix,
	}
}

//func TestMoveFile(t *testing.T) {
//	a := assert.New(t)
//
//	s, err := New("dphoto-unit-move-"+time.Now().Format("20060102150405"), awsSession)
//	must(a, err)
//
//	_, err = s.s3.CreateBucket(&s3.CreateBucketInput{Bucket: &s.bucketName})
//	must(a, err)
//
//	name := "it should copy a file and delete the original"
//
//	_, err = s.UploadFile("unittest", newMedia("skywalker"), "jedi", "anakin.jpg")
//	if a.NoError(err) {
//		_, err = s.MoveFile("unittest", "jedi", "anakin.jpg", "sith")
//		a.NoError(err, name)
//
//		_, err = s.s3.GetObject(&s3.GetObjectInput{
//			Bucket: &s.bucketName,
//			Key:    aws.String("unittest/jedi/anakin.jpg"),
//		})
//		a.Equal(s3.ErrCodeNoSuchKey, err.(awserr.Error).Code(), name)
//
//		_, err = s.s3.GetObject(&s3.GetObjectInput{
//			Bucket: &s.bucketName,
//			Key:    aws.String("unittest/sith/anakin.jpg"),
//		})
//		a.NoError(err, name, name)
//	}
//}

//func must(a *assert.Assertions, err error) {
//	if !a.NoError(err) {
//		a.FailNow(err.Error())
//	}
//}
//
//type InMemoryMedia struct {
//	content string
//}
//
//func (i *InMemoryMedia) SimpleSignature() *backup.SimpleMediaSignature {
//	return &backup.SimpleMediaSignature{
//		RelativePath: "not-used",
//		Size:         uint(len(i.content)),
//	}
//}
//
//func (i *InMemoryMedia) ReadMedia() (io.ReadCloser, error) {
//	return io.NopCloser(strings.NewReader(i.content)), nil
//}
//
//func newMedia(content string) backup.ReadableMedia {
//	return &InMemoryMedia{
//		content: content,
//	}
//}