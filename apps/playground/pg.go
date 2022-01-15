package playground

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/playground/data"
	"gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
	"image/color"
	"strconv"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		th *material.Theme
		ControlPanel
		*apps.Router

		DisableCPBtn widget.Clickable
		isCPDisabled bool

		Tabs
	}

	Tabs struct {
		TabsList []Tab
		list     widget.List
	}

	Tab struct {
		Name       string
		Btn        widget.Clickable
		IsSelected bool
		Layout     func(C, *material.Theme) D
	}
)

func New(router *apps.Router) *Application {
	return &Application{
		Router: router,
		th:     material.NewTheme(gofont.Collection()),
	}
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Tag: &app.DisableCPBtn,
			},
			Layout: func(gtx C, bg, fg color.NRGBA) D {
				var (
					lbl string
				)
				if app.DisableCPBtn.Clicked() {
					app.isCPDisabled = !app.isCPDisabled
				}

				if !app.isCPDisabled {
					lbl = "Disable CP"
				} else {
					lbl = "Enable CP"
				}
				return material.Button(app.th, &app.DisableCPBtn, lbl).Layout(gtx)
			},
		},
	}
}

func (app *Application) Overflow() []component.OverflowAction {
	return []component.OverflowAction{
		{Name: "Close Current Instance - dummy action"},
	}
}

func (app *Application) NavItem() component.NavItem {
	return component.NavItem{
		Name: "playground",
	}
}

func (app *Application) IsCPDisabled() bool {
	return app.isCPDisabled
}

func (app *Application) LayoutView(gtx C, th *material.Theme) D {
	app.initApps()

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return material.List(th, &app.Tabs.list).Layout(gtx, len(app.TabsList), func(gtx C, i int) D {
				var (
					dims D
					btn  material.ButtonStyle
				)
				btn = material.Button(th, &app.TabsList[i].Btn, app.TabsList[i].Name)
				btn.CornerRadius = unit.Dp(1)
				btn.Inset = layout.UniformInset(unit.Dp(10))
				btn.Background = globals.Colours[colours.WHITE]
				btn.Color = globals.Colours[colours.BLACK]
				dims = btn.Layout(gtx)

				if app.TabsList[i].Btn.Clicked() {
					name := app.TabsList[i].Name
					app.TabsList[i].IsSelected = true
					for i := range app.TabsList {
						if name != app.TabsList[i].Name {
							app.TabsList[i].IsSelected = false
						}
					}
					op.InvalidateOp{}.Add(gtx.Ops)
				}

				if app.TabsList[i].IsSelected {
					dims = widget.Border{
						Width:        unit.Dp(1),
						CornerRadius: btn.CornerRadius,
					}.Layout(gtx, func(gtx C) D {
						size := image.Pt(dims.Size.X, dims.Size.Y)
						gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))

						return layout.Stack{Alignment: layout.S}.Layout(gtx,
							layout.Expanded(func(gtx C) D {
								return globals.ColoredArea(gtx, size, globals.Colours[colours.AERO_BLUE])
							}),
							layout.Stacked(func(gtx C) D {
								return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
									layout.Flexed(1, func(gtx C) D {
										var lbl material.LabelStyle
										lbl = material.Body1(th, app.TabsList[i].Name)
										lbl.TextSize = btn.TextSize

										return layout.Flex{}.Layout(gtx,
											layout.Flexed(1, func(gtx C) D {
												return layout.Center.Layout(gtx, lbl.Layout)
											}))
									}),
									layout.Rigid(func(gtx C) D {
										return layout.Stack{}.Layout(gtx,
											layout.Expanded(func(gtx C) D {
												return globals.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, 3), globals.Colours[colours.SEA_GREEN])
											}))
									}))
							}))
					})
				}
				return dims
			})
		}),
		layout.Flexed(1, func(gtx C) D {
			var dims D
			for i := range app.TabsList {
				if app.TabsList[i].IsSelected {
					dims = app.TabsList[i].Layout(gtx, th)
				}
			}
			return dims
		}))
}

func (app *Application) LayoutController(gtx C, th *material.Theme) D {
	return app.ControlPanel.Layout(gtx, th)
}

func (app *Application) initApps() {
	if len(app.TabsList) == 0 {
		app.TabsList = []Tab{

			// Counters app
			{
				Name:       "Counters",
				IsSelected: true,
				Layout: func(gtx C, th *material.Theme) D {
					pgv := data.PgVals

					//TODO find a better way and location to handle the Cache population
					if len(pgv.Cache[data.PRIMES]) == 0 {
						pgv.GenPrimes(data.PLIMIT)
					}
					if len(pgv.Cache[data.FIBS]) == 0 {
						pgv.GenFibs(data.FLIMIT)
					}
					seq := pgv.GetActiveSequence()

					/// DISPLAYED NUMBER
					return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx C) D {
						var val string
						switch seq {
						case data.PRIMES:
							val = strconv.FormatUint(pgv.Cache[seq][pgv.Primes.Index], 10)
						case data.FIBS:
							val = strconv.FormatUint(pgv.Cache[seq][pgv.Fibonacci.Index], 10)
						case data.NATURALS:
							val = strconv.FormatUint(pgv.Naturals.Displayed, 10)
						case data.INTEGERS:
							val = strconv.FormatUint(pgv.Integers.Displayed, 10)
						}
						return material.H5(th, val).Layout(gtx)
					})
				},
			},

			// Draw app
			{
				Name:       "Draw",
				IsSelected: false,
				Layout: func(gtx C, th *material.Theme) D {
					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
							return globals.RColoredArea(gtx, size, 10, globals.Colours[colours.WHITE])
						}),
						layout.Stacked(func(gtx C) D {
							return D{}
						}))
				},
			},
		}
	}
}
