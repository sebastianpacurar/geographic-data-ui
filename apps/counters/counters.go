package counters

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/counters/components/sections"
	"gioui-experiment/custom_themes/colors"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		Top     sections.Top
		View    sections.View
		dockBtn widget.Clickable
		btn     material.IconButtonStyle
		icon    *widget.Icon
		th      *material.Theme
		*apps.Router
	}
)

func New(router *apps.Router) *Application {
	counterApp := &Application{
		Router: router,
	}

	// TODO: add dynamic handling for the diclosers
	//counterApp.View.ControlPanel.InitControllers()
	return counterApp
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Tag: &app.dockBtn,
			},
			Layout: func(gtx C, bg, fg color.NRGBA) D {
				for range app.dockBtn.Clicks() {
					app.NonModalDrawer = !app.NonModalDrawer
				}
				if app.NonModalDrawer {
					app.icon = g.LockCLosedIcon
					app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
					app.btn.Background = bg
					app.btn.Color = g.Colours[colors.DARK_RED]
					app.btn.Size = unit.Dp(24)
				} else {
					app.icon = g.LockOpenedIcon
					app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
					app.btn.Background = bg
					app.btn.Color = g.Colours[colors.SEA_GREEN]
					app.btn.Size = unit.Dp(24)
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
		Name: "Counters - way too overly complicated counter",
	}
}

func (app *Application) Layout(gtx C, th *material.Theme) D {
	top := layout.Rigid(func(gtx C) D {
		return layout.Inset{
			Left:   g.DefaultMargin,
			Right:  g.DefaultMargin,
			Bottom: g.DefaultMargin,
		}.Layout(gtx, func(gtx C) D {
			return app.Top.Layout(gtx, th)
		})
	})
	view := layout.Rigid(func(gtx C) D {
		return app.View.Layout(gtx, th)
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, top, view)
}
