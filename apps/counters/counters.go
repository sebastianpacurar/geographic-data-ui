package counters

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/counters/components/sections"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		Top  sections.Top
		View sections.View
		*apps.Router
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
		Name: "Counters",
	}
}

func (app *Application) Layout(gtx C, th *material.Theme) D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx C) D {
			return g.Inset.Layout(gtx, func(gtx C) D {
				return app.Top.Layout(gtx, th)
			})
		}),
		layout.Flexed(1, func(gtx C) D {
			return app.View.Layout(gtx, th)
		}),
	)
}
