package filter

import (
	"github.com/thomasduchatelle/dphoto/dphoto/backup/backupmodel"
	"github.com/thomasduchatelle/dphoto/dphoto/backup/interactors"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"sync"
)

type filter struct {
	volumeId           string
	lastVolumeSnapshot map[string]uint
	currentSnapshot    []backupmodel.SimpleMediaSignature
	lock               sync.Mutex
}

func NewMediaFilter(volume *backupmodel.VolumeToBackup) (*filter, error) {
	snapshot, err := interactors.VolumeRepositoryPort.RestoreLastSnapshot(volume.UniqueId)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to restore previous snapshot fo volume %s", volume.UniqueId)
	}

	f := &filter{
		volumeId:           volume.UniqueId,
		lastVolumeSnapshot: make(map[string]uint),
		lock:               sync.Mutex{},
	}
	for _, m := range snapshot {
		f.lastVolumeSnapshot[m.RelativePath] = m.Size
	}

	return f, nil
}

func (f *filter) Filter(found backupmodel.FoundMedia) bool {
	f.lock.Lock()
	f.currentSnapshot = append(f.currentSnapshot, *found.SimpleSignature())
	f.lock.Unlock()

	size, ok := f.lastVolumeSnapshot[found.SimpleSignature().RelativePath]
	keep := !ok || size != found.SimpleSignature().Size

	if !keep {
		log.Debugf("Filter > filter out media %s", found)
	}
	return keep
}

func (f *filter) StoreState(backupId string) error {
	return interactors.VolumeRepositoryPort.StoreSnapshot(f.volumeId, backupId, f.currentSnapshot)
}