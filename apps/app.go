package apps

import (
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
	"time"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application interface {
		LayoutView(th *material.Theme) layout.FlexChild
		LayoutController(gtx C, th *material.Theme) D
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
		shutDown       widget.Clickable
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
	for _, e := range r.AppBar.Events(gtx) {
		switch e.(type) {
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
		//r.SetActions([])
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				gtx.Constraints.Max.X = gtx.Px(unit.Dp(250))
				return r.NavDrawer.Layout(gtx, th, &r.NavAnim)
			}),
			layout.Flexed(1, func(gtx C) D {
				size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)

				// view component
				view := layout.Rigid(func(gtx C) D {
					width := gtx.Constraints.Max.X - gtx.Px(g.CountersMenuWidth)
					containerSize := image.Pt(width, gtx.Constraints.Max.Y)
					gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))
					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							container := g.RColoredArea(
								gtx,
								gtx.Constraints.Constrain(size),
								unit.Dp(10),
								g.Colours[colors.ANTIQUE_WHITE],
							)
							return container
						}),
						layout.Stacked(func(gtx C) D {
							containerSize := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
							gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,

								r.pages[r.current].LayoutView(th),
							)
						}))
				})

				// controller component
				controller := layout.Rigid(func(gtx C) D {
					return layout.Stack{Alignment: layout.NW}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.RColoredArea(
								gtx,
								gtx.Constraints.Max,
								unit.Dp(10),
								g.Colours[colors.AERO_BLUE],
							)
						}),
						layout.Stacked(func(gtx C) D {
							containerSize := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
							gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))
							border := widget.Border{
								Color:        g.Colours[colors.SEA_GREEN],
								CornerRadius: unit.Dp(10),
								Width:        unit.Px(1),
							}
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(10),
									Bottom: unit.Dp(10),
								}.Layout(gtx, func(gtx C) D {
									return r.pages[r.current].LayoutController(gtx, th)
								})
							})
						}))
				})

				// lay out the view on the left side and the controller on the right side
				return g.Inset.Layout(gtx, func(gtx C) D {
					return layout.Flex{}.Layout(gtx, view, controller)
				})
			}))
	})
	bar := layout.Rigid(func(gtx C) D {
		return r.AppBar.Layout(gtx, th)
	})

	layout.Flex{Axis: layout.Vertical}.Layout(gtx, bar, content)
	r.ModalLayer.Layout(gtx, th)

	return D{
		Size: gtx.Constraints.Max,
	}
}
