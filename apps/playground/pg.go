package playground

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/playground/tabs"
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

		counterApp tabs.CounterTab
		drawApp    tabs.DrawTab

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

			// Draw app
			{
				Name:       "Draw",
				IsSelected: true,
				Layout: func(gtx C, th *material.Theme) D {
					return app.drawApp.Layout(gtx, th)
				},
			},

			// Counters app
			{
				Name:       "Counters",
				IsSelected: false,
				Layout: func(gtx C, th *material.Theme) D {
					return app.counterApp.Layout(gtx, th)
				},
			},
		}
	}
}
