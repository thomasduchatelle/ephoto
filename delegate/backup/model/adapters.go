package model

import (
	"io"
	"time"
)

type VolumeRepositoryAdapter interface {
	RestoreLastSnapshot(volumeId string) ([]SimpleMediaSignature, error)
	StoreSnapshot(volumeId string, backupId string, signatures []SimpleMediaSignature) error
}

// ClosableMedia can be implemented alongside FoundMedia if the implementation requires to release resources once the media has been handled.
type ClosableMedia interface {
	Close() error
}

type MediaScannerAdapter interface {
	// FindMediaRecursively scan throw the VolumeToBackup and emit to the channel any media found. Interrupted in case of error.
	// returns number of items found, and size of these items
	FindMediaRecursively(volume VolumeToBackup, callback func(FoundMedia)) (uint, uint, error)
}

type ImageDetailsReaderAdapter interface {
	ReadImageDetails(reader io.Reader, lastModifiedDate time.Time) (*MediaDetails, error)
}

type DownloaderAdapter interface {
	DownloadMedia(media FoundMedia) (FoundMedia, error)
}

type OnlineStorageAdapter interface {
	// UploadFile uploads the file in the right folder but might change the name to avoid clash with other existing files. Use files name is always returned.
	UploadFile(media ReadableMedia, folderName, filename string) (string, error)
}

type ReadableMedia interface {
	ReadMedia() (io.Reader, error)
	// SimpleSignature is used to get the size to upload
	SimpleSignature() *SimpleMediaSignature
}