package geography

import (
	"gioui-experiment/apps"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
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
		Name: "Geography - countries, states, statistics",
	}
}

func (app *Application) Layout(gtx C, th *material.Theme) D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			test := material.H2(th, "Geography Application")
			return test.Layout(gtx)
		}),
	)
}
