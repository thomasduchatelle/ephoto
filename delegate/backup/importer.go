package backup

import (
	"duchatelle.io/dphoto/dphoto/backup/backupmodel"
	"duchatelle.io/dphoto/dphoto/backup/interactors"
	"duchatelle.io/dphoto/dphoto/backup/interactors/analyser"
	"github.com/pkg/errors"
	"path"
	"regexp"
	"sort"
	"sync"
)

type scanCompleteListener interface {
	OnScanComplete(total uint)
}

type analyseProgressListener interface {
	OnAnalyseProgress(count, total uint)
}

var (
	datePrefix = regexp.MustCompile("^[0-9]{4}-[01Q][0-9][-_]")
)

// ScanVolume scan a source to discover albums based on original folder structure.
// Listeners will be notified on the progress of the scan.
func ScanVolume(volume backupmodel.VolumeToBackup, listeners ...interface{}) ([]*backupmodel.ScannedFolder, error) {
	medias, err := scanMediaSource(volume)
	if err != nil {
		return nil, err
	}

	triggerScanComplete(listeners, len(medias))

	albums := make(map[string]*backupmodel.ScannedFolder)
	for count, found := range medias {
		_, details, err := analyser.ExtractTypeAndDetails(found)
		if err != nil {
			return nil, err
		}

		dir := path.Dir(found.Filename())
		if album, ok := albums[dir]; ok {
			album.PushBoundaries(details.DateTime, found.SimpleSignature().Size)
		} else {
			albums[dir] = newFoundAlbum(dir)
			albums[dir].PushBoundaries(details.DateTime, found.SimpleSignature().Size)
		}

		triggerProgress(listeners, count, len(medias))
	}

	suggestions := make([]*backupmodel.ScannedFolder, len(albums))
	i := 0
	for _, album := range albums {
		suggestions[i] = album
		i++
	}
	sort.Slice(suggestions, func(i, j int) bool {
		if suggestions[i].Start != suggestions[j].Start {
			return suggestions[i].Start.Before(suggestions[j].Start)
		}

		return suggestions[i].End.Before(suggestions[j].End)
	})

	return suggestions, err
}

func triggerScanComplete(listeners []interface{}, total int) {
	for _, l := range listeners {
		if scanCompleteListener, ok := l.(scanCompleteListener); ok {
			scanCompleteListener.OnScanComplete(uint(total))
		}
	}
}

func triggerProgress(listeners []interface{}, count int, total int) {
	for _, l := range listeners {
		if listener, ok := l.(analyseProgressListener); ok {
			listener.OnAnalyseProgress(uint(count+1), uint(total))
		}
	}
}

func scanMediaSource(volume backupmodel.VolumeToBackup) ([]backupmodel.FoundMedia, error) {
	source, ok := interactors.SourcePorts[volume.Type]
	if !ok {
		return nil, errors.Errorf("No scanner implementation provided for volume type %s", volume.Type)
	}

	lock := sync.Mutex{}
	var medias []backupmodel.FoundMedia
	_, _, err := source.FindMediaRecursively(volume, func(media backupmodel.FoundMedia) {
		lock.Lock()
		medias = append(medias, media)
		lock.Unlock()
	})
	if err != nil {
		return nil, err
	}
	return medias, nil
}

func newFoundAlbum(albumFullPath string) *backupmodel.ScannedFolder {
	name := path.Base(albumFullPath)
	name = datePrefix.ReplaceAllString(name, "")

	return backupmodel.NewScannedFolder(albumFullPath, name)
}
