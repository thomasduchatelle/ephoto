package ui

import (
	"bufio"
	"github.com/buger/goterm"
	screen2 "github.com/thomasduchatelle/ephoto/cmd/fphoto/screen"
	"os"
	"strings"
)

type interactiveRender struct {
	recordsRenderer *recordsRenderer
	screen          *screen2.SimpleScreen
	lastPrintedPage screen2.PagePrint
	form            []*screen2.Segment
	formMode        bool
}

func newInteractiveRender() *interactiveRender {
	return &interactiveRender{
		recordsRenderer: new(recordsRenderer),
		screen:          screen2.NewSimpleScreen(screen2.RenderingOptions{Full: true}),
	}
}

func (i *interactiveRender) Render(state *InteractiveViewState) error {
	table, err := i.recordsRenderer.Render(&state.RecordsState)
	if err != nil {
		return nil
	}

	i.formMode = false
	i.screen.Clear()
	i.lastPrintedPage = screen2.PagePrint{
		Content: []screen2.Segment{screen2.NewConstantSegment(table)},
		Footer:  []screen2.Segment{screen2.NewConstantSegment(strings.Join(state.Actions, " ; "))},
	}
	i.screen.Print(i.lastPrintedPage)

	return nil
}

func (i *interactiveRender) Height() int {
	_, height := i.screen.TermSize()
	return height
}

func (i *interactiveRender) Print(question string) {
	if !i.formMode {
		goterm.MoveCursor(1, i.screen.ContentHeight()+2)
		i.formMode = true
	}
	_, _ = goterm.Print(question)
	goterm.Flush()
}

func (i *interactiveRender) ReadAnswer() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func (i *interactiveRender) TakeOverScreen() {
	i.screen.Clear()
	goterm.MoveCursor(1, 1)
	goterm.Flush()
}
