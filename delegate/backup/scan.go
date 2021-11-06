package backup

import (
	"github.com/thomasduchatelle/dphoto/delegate/backup/backupmodel"
	"github.com/thomasduchatelle/dphoto/delegate/backup/interactors"
	"github.com/thomasduchatelle/dphoto/delegate/backup/interactors/analyser"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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

type ScanOptions struct {
	Listeners   []interface{}
	SkipRejects bool // SkipRejects mode will report any analysis error, or missing timestamp, and continue.
}

var (
	datePrefix = regexp.MustCompile("^[0-9]{4}-[01Q][0-9][-_]")
)

// ScanVolume scan a source to discover albums based on original folder structure.
// Listeners will be notified on the progress of the scan.
func ScanVolume(volume backupmodel.VolumeToBackup, options ScanOptions) ([]*backupmodel.ScannedFolder, []backupmodel.FoundMedia, error) {
	analyser.DefaultMediaTimestamp = !options.SkipRejects

	medias, err := scanMediaSource(volume)
	if err != nil {
		return nil, nil, err
	}

	triggerScanComplete(options.Listeners, len(medias))

	albums := make(map[string]*backupmodel.ScannedFolder)
	var rejects []backupmodel.FoundMedia
	for count, found := range medias {
		_, details, err := analyser.ExtractTypeAndDetails(found)
		if err != nil {
			log.WithError(err).Warnf("Details can't be read from media %s", found)
			rejects = append(rejects, found)

		} else if options.SkipRejects && details.DateTime.IsZero() {
			log.Warnf("Meida timestamp can't be found within the file %s", found)
			rejects = append(rejects, found)

		} else {
			mediaPath := found.MediaPath()
			if album, ok := albums[mediaPath.Path]; ok {
				album.PushBoundaries(details.DateTime, found.SimpleSignature().Size)
			} else {
				albums[mediaPath.Path] = newFoundAlbum(volume, mediaPath)
				albums[mediaPath.Path].PushBoundaries(details.DateTime, found.SimpleSignature().Size)
			}
		}

		triggerProgress(options.Listeners, count, len(medias))
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

	return suggestions, rejects, err
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

func newFoundAlbum(volume backupmodel.VolumeToBackup, mediaPath backupmodel.MediaPath) *backupmodel.ScannedFolder {
	return &backupmodel.ScannedFolder{
		Name:         mediaPath.ParentDir,
		RelativePath: mediaPath.Path,
		FolderName:   mediaPath.ParentDir,
		Distribution: make(map[string]*backupmodel.MediaCounter),
		BackupVolume: &backupmodel.VolumeToBackup{
			UniqueId: volume.UniqueId,
			Type:     volume.Type,
			Path:     mediaPath.ParentFullPath,
			Local:    volume.Local,
		},
	}
}
