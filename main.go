package main

import (
	"gioui-experiment/apps/counters"
	"gioui-experiment/apps/geometry"
	textEditor "gioui-experiment/apps/text_editor/components"
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

const (
	CachePrimes = 200000
	CacheFibs   = 200000
)

var menuBtn = new(widget.Clickable)

func main() {
	ui := newUI()
	go func() {
		w := app.NewWindow(
			app.Title("Gio UI Experiment"),
			app.Size(unit.Dp(1000), unit.Dp(800)),
		)
		if err := ui.Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type UI struct {
	theme      *material.Theme
	navMenu    Menu
	menuItem   MenuItem
	counters   counters.Page
	geometry   geometry.Geometry
	textEditor textEditor.TextEditor
}

type Menu struct {
	active int
	oldVal string
	newVal string
	items  []MenuItem
	list   layout.List
}

type MenuItem struct {
	name       string
	btn        widget.Clickable
	layContent func(gtx C) D
}

func newUI() *UI {
	// Cache is empty only when the app starts up
	globals.CounterVals.GenPrimes(CachePrimes)
	globals.CounterVals.GenFibs(CacheFibs)

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
			name: "Editor",
			layContent: func(gtx C) D {
				return ui.textEditor.Layout(ui.theme, gtx)
			},
		},
	)
	ui.textEditor.InitTextFields()
	ui.counters.Bottom.ValueHandlers.InitTextFields()
	ui.navMenu.oldVal = "Counters"
	return ui
}

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
			// TODO: in progress
			layout.Rigid(func(gtx C) D {
				for i, v := range ui.navMenu.items {
					for range v.btn.Clicks() {
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
									menuItem := &ui.navMenu.items[id]

									// name = the actual name of the application
									// stretches the clickable area to fit the X-Axis
									name := func(gtx C) D {
										return layout.Flex{
											Axis: layout.Horizontal,
										}.Layout(gtx,
											layout.Flexed(1, func(gtx C) D {
												return layout.UniformInset(globals.DefaultMargin).Layout(gtx,
													func(gtx C) D {
														text := material.H6(ui.theme, menuItem.name)
														if gtx.Queue == nil {
															text.Color.A = 150
														}
														return layout.Center.Layout(gtx, text.Layout)
													})
											}))
									}

									// if it's not the current app, then create a clickable area on the whole X-Axis
									if id != ui.navMenu.active {
										return material.Clickable(gtx, &menuItem.btn, func(gtx C) D {
											return name(gtx)
										})
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
