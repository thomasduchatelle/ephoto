package daemon

import (
	"duchatelle.io/dphoto/dphoto/backup/backupmodel"
)

var (
	VolumeManager VolumeManagerPort
)

type VolumeManagerPort interface {
	OnMountedVolume(volume backupmodel.VolumeToBackup)
	OnUnMountedVolume(uuid string)
}
