package geography

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/geography/components"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
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
		dockBtn widget.Clickable
		icon    *widget.Icon
		btn     material.IconButtonStyle
		th      *material.Theme
		components.Geography
		components.ControlPanel
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
				} else {
					app.icon = g.LockOpenedIcon
					app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
					app.btn.Background = bg
					app.btn.Color = g.Colours[colors.SEA_GREEN]
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
		Name: "Geography - countries, states, statistics",
	}
}

func (app *Application) LayoutView(th *material.Theme) layout.FlexChild {
	return app.Geography.Layout(th)
}

func (app *Application) LayoutController(gtx C, th *material.Theme) D {
	return app.ControlPanel.Layout(gtx, th)
}
