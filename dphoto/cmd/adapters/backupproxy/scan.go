package backupproxy

import (
	"crypto/sha256"
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"github.com/pkg/errors"
	"github.com/thomasduchatelle/dphoto/domain/backup"
	"github.com/thomasduchatelle/dphoto/dphoto/cmd/ui"
	"github.com/thomasduchatelle/dphoto/dphoto/screen"
)

type ScanProgress struct {
	screen           *screen.AutoRefreshScreen
	scanningLine     *screen.ProgressLine
	analysedLine     *screen.ProgressLine
	onAnalysedCalled bool
}

func ScanWithCache(owner string, volume backup.SourceVolume, options ...backup.Options) (ui.SuggestionRecordRepositoryPort, []backup.FoundMedia, error) {
	previousResult, rejectCount, err := restore(volumeId(volume))
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to restore previous scan result for volume %s", volume.String())
	}

	if len(previousResult) > 0 {
		useState, ok := ui.NewSimpleForm().ReadBool("Previous result has been found for this volume, do you want to restore it?", "Y/n")
		if !ok || useState {
			return NewSuggestionRepository(owner, previousResult, rejectCount), nil, err
		}
	}

	suggestions, rejects, err := doScan(owner, volume, options...)
	return NewSuggestionRepository(owner, suggestions, len(rejects)), rejects, err
}

func volumeId(volume backup.SourceVolume) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(volume.String())))
}

func doScan(owner string, volume backup.SourceVolume, options ...backup.Options) ([]*backup.ScannedFolder, []backup.FoundMedia, error) {
	progress := newScanProgress()
	options = append(options, backup.OptionWithListener(progress))

	suggestions, rejects, err := backup.Scan(owner, volume, options...)
	progress.stop()

	if err != nil {
		return nil, nil, err
	}

	err = Store(volumeId(volume), suggestions, len(rejects))
	return suggestions, rejects, err
}

func newScanProgress() *ScanProgress {
	table := screen.NewTable(" ", 2, 20, 80, 25)
	scanningLine, scanningSegment := screen.NewProgressLine(table, "Scanning...")
	analysedLine, analysedSegment := screen.NewProgressLine(table, "Analysing...")

	progressScreen := screen.NewAutoRefreshScreen(
		screen.RenderingOptions{Width: 180},
		scanningSegment,
		analysedSegment,
	)

	return &ScanProgress{
		screen:       progressScreen,
		scanningLine: scanningLine,
		analysedLine: analysedLine,
	}
}

func (s *ScanProgress) OnScanComplete(total backup.MediaCounter) {
	s.scanningLine.SwapSpinner(1)
	s.scanningLine.SetLabel(aurora.Sprintf("%d files has been found.", aurora.Cyan(total.Count)))
}

func (s *ScanProgress) OnAnalysed(done, total backup.MediaCounter) {
	s.onAnalysedCalled = true
	if !total.IsZero() {
		s.analysedLine.SetBar(done.Count, total.Count)
		s.analysedLine.SetExplanation(fmt.Sprintf("%d / %d files", done.Count, total.Count))

		if done.Count == total.Count {
			s.analysedLine.SwapSpinner(1)
			s.analysedLine.SetLabel("Analyse complete")
		}
	}
}

func (s *ScanProgress) stop() {
	if !s.onAnalysedCalled {
		s.analysedLine.SwapSpinner(1)
		s.analysedLine.SetLabel("Analyse skipped")
	}
	s.screen.Stop()
}
