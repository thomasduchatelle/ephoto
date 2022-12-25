package backupui

import (
	"fmt"
	screen2 "github.com/thomasduchatelle/ephoto/cmd/dphoto/screen"
	"github.com/thomasduchatelle/ephoto/pkg/backup"
)

type BackupProgress struct {
	screen           *screen2.AutoRefreshScreen
	scanLine         *screen2.ProgressLine
	analyseLine      *screen2.ProgressLine
	uploadLine       *screen2.ProgressLine
	onAnalysedCalled bool
	onUploadCalled   bool
}

func NewProgress() *BackupProgress {
	table := screen2.NewTable(" ", 2, 20, 80, 25)

	segments := make([]screen2.Segment, 3)
	p := &BackupProgress{}
	p.scanLine, segments[0] = screen2.NewProgressLine(table, "Scanning...")
	p.analyseLine, segments[1] = screen2.NewProgressLine(table, "Analysing...")
	p.uploadLine, segments[2] = screen2.NewProgressLine(table, "Uploading ...")

	p.screen = screen2.NewAutoRefreshScreen(
		screen2.RenderingOptions{Width: 180},
		segments...,
	)

	return p
}

func (p *BackupProgress) OnScanComplete(total backup.MediaCounter) {
	if total.Count == 0 {
		p.scanLine.SwapSpinner(1)
		p.scanLine.SetLabel(fmt.Sprintf("Scan complete: no new files found"))
	} else {
		p.scanLine.SwapSpinner(1)
		p.scanLine.SetLabel(fmt.Sprintf("Scan complete: %d files found", total.Count))
	}
}

func (p *BackupProgress) OnAnalysed(done, total backup.MediaCounter) {
	p.onAnalysedCalled = true
	if !total.IsZero() {
		p.analyseLine.SetBar(done.Count, total.Count)
		p.analyseLine.SetExplanation(fmt.Sprintf("%d / %d files", done.Count, total.Count))

		if done.Count == total.Count {
			p.analyseLine.SwapSpinner(1)
			p.analyseLine.SetLabel("Analyse complete")
		}
	}
}

func (p *BackupProgress) OnUploaded(done, total backup.MediaCounter) {
	p.onUploadCalled = true
	if !total.IsZero() {
		p.uploadLine.SetBar(done.Size, total.Size)
		p.uploadLine.SetExplanation(fmt.Sprintf("%s / %s", byteCountIEC(done.Size), byteCountIEC(total.Size)))

		if done.Count == total.Count {
			p.uploadLine.SwapSpinner(1)
			p.uploadLine.SetLabel("Upload complete")
		}
	} else {
		p.uploadLine.SetExplanation(fmt.Sprintf("%s", byteCountIEC(done.Size)))
	}
}

func (p *BackupProgress) Stop() {
	if !p.onAnalysedCalled {
		p.analyseLine.SwapSpinner(1)
		p.analyseLine.SetLabel("Analyse skipped")
	}
	if !p.onUploadCalled {
		p.uploadLine.SwapSpinner(1)
		p.uploadLine.SetLabel("Upload skipped")
	}
	p.screen.Stop()
}
