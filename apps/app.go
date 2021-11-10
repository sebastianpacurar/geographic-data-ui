package apps

import (
	g "gioui-experiment/globals"
	"gioui.org/layout"
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
	bar.NavigationIcon = g.MenuIcon

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

	content := layout.Flexed(1, func(gtx C) D {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				gtx.Constraints.Max.X /= 4
				return r.NavDrawer.Layout(gtx, th, &r.NavAnim)
			}),
			layout.Flexed(1, func(gtx C) D {
				return g.Inset.Layout(gtx, func(gtx C) D {
					return r.pages[r.current].Layout(gtx, th)
				})
			}),
		)
	})
	bar := layout.Rigid(func(gtx C) D {
		return r.AppBar.Layout(gtx, th)
	})

	// lay the app bar first, then the content of the current application
	layout.Flex{Axis: layout.Vertical}.Layout(gtx, bar, content)

	// lay the modal on top of other widgets, so it could have the highest z-index
	r.ModalLayer.Layout(gtx, th)

	return D{
		Size: gtx.Constraints.Max,
	}
}
