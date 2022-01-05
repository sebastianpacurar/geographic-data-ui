package playground

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/playground/data"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
	"strconv"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		//dockBtn widget.Clickable
		//btn     material.IconButtonStyle
		//icon    *widget.Icon
		th *material.Theme
		ControlPanel
		*apps.Router

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
	}
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		// TODO: think what to do with this only on PG
		//{
		//	OverflowAction: component.OverflowAction{
		//		Tag: &app.dockBtn,
		//	},
		//	Layout: func(gtx C, bg, fg color.NRGBA) D {
		//		for range app.dockBtn.Clicks() {
		//			app.NonModalDrawer = !app.NonModalDrawer
		//		}
		//		if app.NonModalDrawer {
		//			app.icon = g.LockCLosedIcon
		//			app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
		//			app.btn.Background = bg
		//			app.btn.Color = g.Colours[colours.DARK_RED]
		//			app.btn.Size = unit.Dp(24)
		//		} else {
		//			app.icon = g.LockOpenedIcon
		//			app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
		//			app.btn.Background = bg
		//			app.btn.Color = g.Colours[colours.SEA_GREEN]
		//			app.btn.Size = unit.Dp(24)
		//		}
		//		return app.btn.Layout(gtx)
		//	},
		//},
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
				btn.Background = g.Colours[colours.WHITE]
				btn.Color = g.Colours[colours.BLACK]
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
						Width:        unit.Px(1),
						CornerRadius: unit.Dp(1),
					}.Layout(gtx, func(gtx C) D {
						size := image.Pt(dims.Size.X, dims.Size.Y)
						return layout.Stack{}.Layout(gtx,
							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, size, g.Colours[colours.AERO_BLUE])
							}),
							layout.Stacked(func(gtx C) D {
								return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx C) D {
									var lbl material.LabelStyle
									lbl = material.Body1(th, app.TabsList[i].Name)
									lbl.TextSize = btn.TextSize
									return lbl.Layout(gtx)
								})
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
					return layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(50),
						Bottom: unit.Dp(20),
						Left:   unit.Dp(50),
					}.Layout(gtx, func(gtx C) D {
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
					return D{}
				},
			},
		}
	}
}
