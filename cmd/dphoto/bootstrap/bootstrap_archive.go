package bootstrap

import (
	config2 "github.com/thomasduchatelle/ephoto/cmd/dphoto/config"
	"github.com/thomasduchatelle/ephoto/pkg/archive"
	"github.com/thomasduchatelle/ephoto/pkg/archiveadapters/archivedynamo"
	"github.com/thomasduchatelle/ephoto/pkg/archiveadapters/asyncjobadapter"
	"github.com/thomasduchatelle/ephoto/pkg/archiveadapters/s3store"
)

func init() {
	config2.Listen(func(cfg config2.Config) {
		repositoryAdapter := archivedynamo.Must(archivedynamo.New(cfg.GetAWSSession(), cfg.GetString(config2.ArchiveDynamodbTable), false))
		storeAdapter := s3store.Must(s3store.New(cfg.GetAWSSession(), cfg.GetString(config2.ArchiveMainBucketName)))
		cacheAdapter := s3store.Must(s3store.New(cfg.GetAWSSession(), cfg.GetString(config2.ArchiveCacheBucketName)))
		archiveAsyncAdapter := asyncjobadapter.New(cfg.GetAWSSession(), cfg.GetString(config2.ArchiveJobsSNSARN), cfg.GetString(config2.ArchiveJobsSQSURL), asyncjobadapter.DefaultImagesPerMessage)
		archive.Init(
			repositoryAdapter,
			storeAdapter,
			cacheAdapter,
			archiveAsyncAdapter,
		)
	})
}
