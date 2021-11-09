package editor

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/editor/components"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		editor components.TextArea
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
		Name: "Editor - multifunctional text editor",
	}
}

func (app *Application) Layout(gtx C, th *material.Theme) D {
	return app.editor.Layout(gtx, th)
}
