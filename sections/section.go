package sections

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
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

	Section interface {
		LayoutView(gtx C, th *material.Theme) D
		LayoutController(gtx C, th *material.Theme) D
		Actions() []component.AppBarAction
		Overflow() []component.OverflowAction
		NavItem() component.NavItem
		IsCPDisabled() bool
	}

	Router struct {
		pages   map[interface{}]Section
		current interface{}
		*component.ModalNavDrawer
		NavAnim component.VisibilityAnimation
		*component.AppBar
		*component.ModalLayer
		component.Resize
	}
)

func NewRouter() Router {
	modal := component.NewModal()

	nav := component.NewNav("Gioui Experiment", "Available Sections")
	modalNav := component.ModalNavFrom(&nav, modal)

	bar := component.NewAppBar(modal)
	bar.NavigationIcon = globals.MenuIcon

	na := component.VisibilityAnimation{
		State:    component.Invisible,
		Duration: time.Millisecond * 250,
	}
	return Router{
		pages:          make(map[interface{}]Section),
		ModalLayer:     modal,
		ModalNavDrawer: modalNav,
		AppBar:         bar,
		NavAnim:        na,
		Resize:         component.Resize{Ratio: 0.70},
	}
}

func (r *Router) Register(tag interface{}, app Section) {
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
			r.ModalNavDrawer.Appear(gtx.Now)
			r.NavAnim.Disappear(gtx.Now)
		}
	}

	if r.ModalNavDrawer.NavDestinationChanged() {
		r.SwitchTo(r.ModalNavDrawer.CurrentNavDestination())
	}

	content := layout.Flexed(1, func(gtx C) D {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				gtx.Constraints.Max.X = gtx.Px(unit.Dp(250))
				return r.NavDrawer.Layout(gtx, th, &r.NavAnim)
			}),
			layout.Flexed(1, func(gtx C) D {
				var dims D
				if r.pages[r.current].IsCPDisabled() {

					dims = layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							container := globals.ColoredArea(
								gtx,
								image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y),
								globals.Colours[globals.ANTIQUE_WHITE],
							)
							return container
						}),
						layout.Stacked(func(gtx C) D {
							containerSize := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
							gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								layout.Rigid(func(gtx C) D {
									return layout.Inset{
										Right:  unit.Dp(10),
										Bottom: unit.Dp(5),
										Left:   unit.Dp(10),
									}.Layout(gtx, func(gtx C) D {
										return r.pages[r.current].LayoutView(gtx, th)
									})
								}))
						}))
				} else {
					dims = r.Resize.Layout(gtx,
						func(gtx C) D {
							return layout.Stack{}.Layout(gtx,
								layout.Expanded(func(gtx C) D {
									container := globals.ColoredArea(
										gtx,
										gtx.Constraints.Max,
										globals.Colours[globals.ANTIQUE_WHITE],
									)
									return container
								}),
								layout.Stacked(func(gtx C) D {
									return layout.Inset{
										Right:  unit.Dp(10),
										Bottom: unit.Dp(5),
										Left:   unit.Dp(10),
									}.Layout(gtx, func(gtx C) D {
										return r.pages[r.current].LayoutView(gtx, th)
									})
								}))
						},
						func(gtx C) D {
							return layout.Stack{Alignment: layout.NW}.Layout(gtx,
								layout.Expanded(func(gtx C) D {
									return globals.ColoredArea(
										gtx,
										gtx.Constraints.Max,
										globals.Colours[globals.AERO_BLUE],
									)
								}),
								layout.Stacked(func(gtx C) D {
									containerSize := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
									gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))
									border := widget.Border{
										Color: globals.Colours[globals.SEA_GREEN],
										Width: unit.Dp(1),
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
						},
						func(gtx C) D {
							rect := image.Rectangle{
								Max: image.Point{
									X: gtx.Px(unit.Dp(6)),
									Y: gtx.Constraints.Max.Y,
								},
							}
							paint.FillShape(gtx.Ops, globals.Colours[globals.CP_RESIZER], clip.Rect(rect).Op())
							return D{Size: rect.Max}
						})
				}
				return dims
			}))
	})

	bar := layout.Rigid(func(gtx C) D {
		return r.AppBar.Layout(gtx, th, "desc", "desc")
	})

	layout.Flex{Axis: layout.Vertical}.Layout(gtx, bar, content)
	r.ModalLayer.Layout(gtx, th)

	return D{
		Size: gtx.Constraints.Max,
	}
}
