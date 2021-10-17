package main

import (
	"gioui-experiment/app_layout"
	counters "gioui-experiment/apps/counters"
	formatters "gioui-experiment/apps/formatters/components"
	"gioui-experiment/apps/geometry"
	"gioui-experiment/globals"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"log"
	"os"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	ui := newUI()

	// Starts a goroutine which executes an anonymous function.
	// Starts the app and initializes the UI.
	go func() {
		w := app.NewWindow(
			app.Title("Gio UI Experiment"),
			app.Size(unit.Dp(1000), unit.Dp(800)),
		)

		// if err -> os.Exit(1), quits the goroutine
		if err := ui.Run(w); err != nil {
			log.Fatal(err)
		}

		// 0 is fine so the goroutine continues
		os.Exit(0)
	}()
	app.Main()
}

// UI holds the entire states of the app.
type UI struct {
	theme         *material.Theme
	topBar        app_layout.TopBar
	navMenu       app_layout.Menu
	counters      counters.Page
	geometry      geometry.Geometry
	jsonFormatter formatters.JsonFormatter
}

// newUI returns a new UI which uses the Go Fonts, and initializes the Text Fields states
func newUI() *UI {
	ui := &UI{
		theme: material.NewTheme(gofont.Collection()),
	}
	//ui.menu.Items = append(ui.menu.Items,
	//	Menu{
	//		name: "Counters",
	//		w: func(item *menuItem, gtx C) D {
	//			return outl
	//		},
	//	})
	ui.jsonFormatter.InitTextFields()
	ui.counters.TopController.InitTextFields()
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
			// Reset the layout Context for a new frame.
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

// Layout - displays the content of the application.
// Inset refers to the margins of the components, so there can be
// a small margin around the entire contents of the app.
func (ui *UI) Layout(gtx C) D {
	windowBorder := widget.Border{
		Color:        globals.Colours["dark-cyan"],
		CornerRadius: unit.Dp(0),
		Width:        unit.Dp(3),
	}
	return windowBorder.Layout(gtx, func(gtx C) D {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(
			gtx,
			layout.Rigid(func(gtx C) D {
				return ui.topBar.Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(
					gtx,
					layout.Rigid(func(gtx C) D {
						return ui.navMenu.Layout(ui.theme, gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return ui.counters.Layout(ui.theme, gtx)
					}),
				)
			}),
		)
	},
	)
}
