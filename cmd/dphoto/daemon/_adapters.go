package daemon

import (
	"github.com/thomasduchatelle/ephoto/dphoto/backup/backupmodel"
)

var (
	VolumeManager VolumeManagerPort
)

type VolumeManagerPort interface {
	OnMountedVolume(volume backupmodel.VolumeToBackup)
	OnUnMountedVolume(uuid string)
}
