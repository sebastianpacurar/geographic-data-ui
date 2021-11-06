package apps

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"time"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application interface {
		Layout(gtx C, th *material.Theme) D
		Actions() []component.AppBarAction
		Overflow() []component.OverflowAction
		NavItem() component.NavItem
	}

	Router struct {
		pages   map[interface{}]Application
		current interface{}
		*component.ModalNavDrawer
		NavAnim component.VisibilityAnimation
		*component.AppBar
		*component.ModalLayer
		NonModalDrawer bool
	}
)

func NewRouter() Router {
	modal := component.NewModal()

	nav := component.NewNav("Gioui Experiment", "Click or Tap on any of the below apps")
	modalNav := component.ModalNavFrom(&nav, modal)

	bar := component.NewAppBar(modal)
	bar.NavigationIcon = globals.MenuIcon

	na := component.VisibilityAnimation{
		State:    component.Invisible,
		Duration: time.Millisecond * 250,
	}
	return Router{
		pages:          make(map[interface{}]Application),
		ModalLayer:     modal,
		ModalNavDrawer: modalNav,
		AppBar:         bar,
		NavAnim:        na,
	}
}

func (r *Router) Register(tag interface{}, app Application) {
	r.pages[tag] = app
	navItem := app.NavItem()
	navItem.Tag = tag
	if r.current == interface{}(nil) {
		r.current = tag
		r.AppBar.Title = navItem.Name
		r.AppBar.SetActions(app.Actions(), app.Overflow())
	}
	r.ModalNavDrawer.AddNavItem(navItem)
}

func (r *Router) SwitchTo(tag interface{}) {
	app, ok := r.pages[tag]
	if !ok {
		return
	}
	navItem := app.NavItem()
	r.current = tag
	r.AppBar.Title = navItem.Name
	r.AppBar.SetActions(app.Actions(), app.Overflow())
}

func (r *Router) Layout(gtx C, th *material.Theme) D {
	for _, event := range r.AppBar.Events(gtx) {
		switch event.(type) {
		case component.AppBarNavigationClicked:
			if r.NonModalDrawer {
				r.NavAnim.ToggleVisibility(gtx.Now)
			} else {
				r.ModalNavDrawer.Appear(gtx.Now)
				r.NavAnim.Disappear(gtx.Now)
			}
		}
	}
	if r.ModalNavDrawer.NavDestinationChanged() {
		r.SwitchTo(r.ModalNavDrawer.CurrentNavDestination())
	}

	paint.Fill(gtx.Ops, th.Palette.Bg)
	content := layout.Flexed(1, func(gtx C) D {
		return layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				gtx.Constraints.Max.X /= 3
				return r.NavDrawer.Layout(gtx, th, &r.NavAnim)
			}),
			layout.Flexed(1, func(gtx C) D {
				return r.pages[r.current].Layout(gtx, th)
			}),
		)
	})
	bar := layout.Rigid(func(gtx C) D {
		return r.AppBar.Layout(gtx, th)
	})
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx, bar, content)

	r.ModalLayer.Layout(gtx, th)
	return D{
		Size: gtx.Constraints.Max,
	}
}
