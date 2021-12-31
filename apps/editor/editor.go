package editor

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/editor/components"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		dockBtn widget.Clickable
		icon    *widget.Icon
		btn     material.IconButtonStyle
		th      *material.Theme
		components.TextArea
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

func (app *Application) LayoutView(th *material.Theme) layout.FlexChild {
	return app.TextArea.Layout(th)
}

func (app *Application) LayoutController(gtx C, th *material.Theme) D {
	return app.ControlPanel.Layout(gtx, th)
}
