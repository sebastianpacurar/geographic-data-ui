package main

import (
	comp "gioui-experiment/multioption_counter/components"
	"gioui-experiment/multioption_counter/globals"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"log"
	"os"
)

func main() {
	ui := NewUI()
	go func() {
		w := app.NewWindow(
			app.Title("Multi Option Counter"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)
		if err := ui.Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

// UI holds the entire states of the app.
type UI struct {
	theme      *material.Theme
	counter    comp.Counter
	startValue comp.StartValue
}

// NewUI returns a new UI which uses the Go Fonts.
func NewUI() *UI {
	ui := &UI{}
	ui.theme = material.NewTheme(gofont.Collection())
	return ui
}

// Run renders the application and responds to different events.
// ops are the operations passed to the graphics context (gtx)
// system.FrameEvent - this is sent when the application receives a re-render event:
// it sets the context with the operations and the event. this is used to pass
// around event information.
// key.NameEscape - returning null means shut down the application.
// system.DestroyEvent - this is sent when the application closes.
func (ui *UI) Run(w *app.Window) error {
	var ops op.Ops
	for event := range w.Events() {
		switch event := event.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, event)
			ui.Layout(gtx)
			event.Frame(gtx.Ops)
		case key.Event:
			switch event.Name {
			case key.NameEscape:
				return nil
			}
		case system.DestroyEvent:
			return event.Err
		}
	}
	return nil
}

// Layout - displays the startValue and counter components vertically.
// Inset refers to the margins of the components, so there can be
// a small margin around the entire contents of the app.
func (ui *UI) Layout(gtx globals.C) globals.D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx globals.C) globals.D {
			return globals.Inset.Layout(gtx, func(gtx globals.C) globals.D {
				return ui.startValue.Layout(ui.theme, gtx)
			})
		}),
		layout.Rigid(func(gtx globals.C) globals.D {
			return globals.Inset.Layout(gtx, func(gtx globals.C) globals.D {
				return ui.counter.Layout(ui.theme, gtx)
			})
		}),
	)
}