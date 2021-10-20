package main

import (
	"gioui-experiment/app_layout"
	"gioui-experiment/apps/counters"
	formatters "gioui-experiment/apps/formatters/components"
	"gioui-experiment/apps/geometry"
	"gioui-experiment/globals"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
	"image/color"
	"log"
	"os"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

var (
	menuBtn = new(widget.Clickable)
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
	navMenu       Menu
	menuItem      MenuItem
	counters      counters.Page
	geometry      geometry.Geometry
	jsonFormatter formatters.JsonFormatter
}

type Menu struct {
	active int
	items  []MenuItem
	list   layout.List
}

type MenuItem struct {
	name       string
	btn        widget.Clickable
	layContent func(gtx C) D
}

// newUI returns a new UI which uses the Go Fonts, and initializes the Text Fields states
func newUI() *UI {
	ui := &UI{
		theme: material.NewTheme(gofont.Collection()),
	}
	ui.navMenu.list.Axis = layout.Vertical
	ui.navMenu.items = append(ui.navMenu.items,
		MenuItem{
			name: "Counters",
			layContent: func(gtx C) D {
				return ui.counters.Layout(ui.theme, gtx)
			},
		},
		MenuItem{
			name: "Formatters",
			layContent: func(gtx C) D {
				return ui.jsonFormatter.Layout(ui.theme, gtx)
			},
		},
	)
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

	// activeApp = the currently selected menu item
	activeApp := &ui.navMenu.items[ui.navMenu.active]
	return windowBorder.Layout(gtx, func(gtx C) D {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(
			gtx,

			/// TOP BAR SECTION
			layout.Rigid(func(gtx C) D {
				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx C) D {
						size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y+10/2)
						bar := globals.ColoredArea(
							gtx,
							gtx.Constraints.Constrain(size),
							globals.Colours["dark-cyan"],
						)
						return bar
					}),
					layout.Stacked(func(gtx C) D {
						return layout.Inset{
							Left: unit.Dp(10),
							Top:  unit.Dp(5),
						}.Layout(gtx, func(gtx C) D {
							btn := component.SimpleIconButton(
								globals.Colours["dark-cyan"],
								globals.Colours["white"],
								menuBtn,
								globals.MenuIcon,
							)
							return btn.Layout(gtx)
						})
					}),
				)
			}),

			// NAVIGATION MENU SECTION
			layout.Rigid(func(gtx C) D {
				for i := range ui.navMenu.items {
					for ui.navMenu.items[i].btn.Clicked() {
						ui.navMenu.active = i
					}
				}

				return layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(
					gtx,
					layout.Rigid(func(gtx C) D {
						width := gtx.Px(globals.MenuWidth)
						containerSize := image.Pt(width, gtx.Constraints.Max.Y)
						gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))

						return layout.Stack{
							Alignment: layout.NW,
						}.Layout(
							gtx,
							layout.Expanded(func(gtx C) D {
								size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
								bar := globals.ColoredArea(
									gtx,
									gtx.Constraints.Constrain(size),
									globals.Colours["sea-green"],
								)
								return bar
							}),

							layout.Stacked(func(gtx C) D {
								return ui.navMenu.list.Layout(gtx, len(ui.navMenu.items), func(gtx C, id int) D {
									application := &ui.navMenu.items[id]

									// name = the actual name of the application
									// stretches the clickable area to fit the X-Axis
									name := func(gtx C) D {
										return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
											layout.Flexed(1, func(gtx C) D {
												return layout.UniformInset(globals.DefaultMargin).Layout(gtx, material.H6(ui.theme, application.name).Layout)
											}))
									}

									// if it's not the current app, then create a clickable area on the whole X-Axis
									if id != ui.navMenu.active {
										return material.Clickable(gtx, &application.btn, name)
									}

									// lay out the selected item in a grey-ish background
									return layout.Stack{}.Layout(gtx,
										layout.Expanded(func(gtx C) D {
											clip.UniformRRect(f32.Rectangle{
												Max: layout.FPt(gtx.Constraints.Max),
											}, 0).Add(gtx.Ops)
											paint.Fill(gtx.Ops, color.NRGBA{A: 64})
											return D{}
										}),
										layout.Stacked(name),
									)
								})
							}),
						)
					}),

					/// APPLICATION SECTION
					layout.Rigid(func(gtx C) D {
						return activeApp.layContent(gtx)
					}))
			}),
		)
	})
}
