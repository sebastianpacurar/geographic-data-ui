package editor

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/editor/components"
	"gioui-experiment/custom_themes/colors"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		editor  components.TextArea
		dockBtn widget.Clickable
		icon    *widget.Icon
		btn     material.IconButtonStyle
		th      *material.Theme
		*apps.Router
	}
)

func New(router *apps.Router) *Application {
	return &Application{
		Router: router,
	}
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Tag: &app.dockBtn,
			},
			Layout: func(gtx C, bg, fg color.NRGBA) D {
				if app.dockBtn.Clicked() {
					app.NonModalDrawer = !app.NonModalDrawer
				}
				if app.NonModalDrawer {
					app.icon = g.LockCLosedIcon
					app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
					app.btn.Background = bg
					app.btn.Color = g.Colours[colors.DARK_RED]
					app.Router.NavAnim.Appear(gtx.Now)
				} else {
					app.icon = g.LockOpenedIcon
					app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
					app.btn.Background = bg
					app.btn.Color = g.Colours[colors.SEA_GREEN]
					app.Router.NavAnim.Disappear(gtx.Now)
				}
				return app.btn.Layout(gtx)
			},
		},
	}
}

func (app *Application) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (app *Application) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Editor - multifunctional text editor",
	}
}

func (app *Application) Layout(gtx C, th *material.Theme) D {
	return app.editor.Layout(gtx, th)
}
