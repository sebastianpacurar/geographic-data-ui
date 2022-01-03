package editor

import (
	"gioui-experiment/apps"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		th *material.Theme
		TextArea
		ControlPanel
		*apps.Router
	}

	TextArea struct {
		Field widget.Editor
		List  widget.List

		Menu            component.MenuState
		CtxArea         component.ContextArea
		IsMenuTriggered bool

		// menu options
		PasteBtn widget.Clickable
	}
)

func New(router *apps.Router) *Application {
	return &Application{
		Router: router,
	}
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (app *Application) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (app *Application) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Editor - multifunctional text editor",
	}
}

func (app *Application) LayoutView(gtx C, th *material.Theme) D {
	app.TextArea.List.Axis = layout.Vertical
	app.TextArea.Field.SingleLine = false
	app.TextArea.Field.Alignment = text.Start

	if !app.TextArea.IsMenuTriggered {
		var item component.MenuItemStyle
		item.LabelInset = layout.Inset{
			Top:    unit.Dp(5),
			Right:  unit.Dp(5),
			Bottom: unit.Dp(5),
			Left:   unit.Dp(5),
		}
		item = component.MenuItem(th, &app.TextArea.PasteBtn, "Paste")
		app.TextArea.Menu = component.MenuState{
			Options: []func(gtx C) D{
				item.Layout,
			},
		}
	}

	border := widget.Border{
		Color:        g.Colours[colors.GREY],
		CornerRadius: unit.Dp(5),
		Width:        unit.Px(2),
	}
	switch {
	case app.TextArea.Field.Focused():
		border.Color = th.Palette.ContrastBg
		border.Width = unit.Px(2)
	}

	return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx C) D {
		return layout.Stack{}.Layout(gtx,
			layout.Stacked(func(gtx C) D {
				gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(gtx.Constraints.Max))
				return border.Layout(gtx, func(gtx C) D {
					return material.List(th, &app.TextArea.List).Layout(gtx, 1, func(gtx C, _ int) D {
						return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
							ed := material.Editor(th, &app.TextArea.Field, "Type your Thoughts...")
							ed.SelectionColor = g.Colours[colors.TEXT_SELECTION]

							if app.TextArea.PasteBtn.Clicked() {
								ed.Editor.SetText(g.ClipBoardVal)
							}
							return ed.Layout(gtx)
						})
					})
				})
			}),
			layout.Expanded(func(gtx C) D {
				return app.TextArea.CtxArea.Layout(gtx, func(gtx C) D {
					gtx.Constraints.Min = image.Point{}
					return component.Menu(th, &app.TextArea.Menu).Layout(gtx)
				})
			}))
	})
}

func (app *Application) LayoutController(gtx C, th *material.Theme) D {
	return app.ControlPanel.Layout(gtx, th)
}
