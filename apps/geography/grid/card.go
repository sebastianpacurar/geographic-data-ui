package grid

import (
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
)

type (
	Card struct {
		Name string
		Flag image.Image

		Active   bool
		Selected bool

		menu            component.MenuState
		ctxArea         component.ContextArea
		isMenuTriggered bool

		// menu options
		selectBtn     widget.Clickable
		deselectBtn   widget.Clickable
		copyToClipBtn widget.Clickable
		viewBtn       widget.Clickable
	}
)

func (c *Card) LayCard(gtx C, th *material.Theme) D {
	size := image.Pt(200, 250)
	var container component.SurfaceStyle

	container.Theme = material.NewTheme(gofont.Collection())
	container.Theme.Bg = g.Colours[colours.CARD_COLOR]
	container.Elevation = unit.Dp(5)
	container.ShadowStyle = component.Shadow(unit.Dp(5), unit.Dp(8))

	if !c.isMenuTriggered {
		lbl := "Select"
		btn := &c.selectBtn
		if c.Selected {
			lbl = "Deselect"
			btn = &c.deselectBtn
		}
		var item component.MenuItemStyle
		item.LabelInset = layout.Inset{
			Top:    unit.Dp(5),
			Right:  unit.Dp(5),
			Bottom: unit.Dp(5),
			Left:   unit.Dp(5),
		}
		item = component.MenuItem(th, btn, lbl)

		c.menu = component.MenuState{
			Options: []func(gtx C) D{
				item.Layout,
				component.MenuItem(th, &c.copyToClipBtn, "Copy as JSON").Layout,
				component.MenuItem(th, &c.viewBtn, "View").Layout,
			},
		}
	}
	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx C) D {
			return container.Layout(gtx, func(gtx C) D {
				gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx C) D {
					//gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))
					return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceAround}.Layout(gtx,
						// country name
						layout.Rigid(func(gtx C) D {
							return layout.Flex{}.Layout(gtx,
								layout.Flexed(1, func(gtx C) D {
									return layout.Center.Layout(gtx, func(gtx C) D {
										return material.Body2(th, c.Name).Layout(gtx)
									})
								}),
							)
						}),

						// country flag image
						layout.Rigid(func(gtx C) D {
							return layout.Flex{}.Layout(gtx,
								layout.Flexed(1, func(gtx C) D {
									return layout.Center.Layout(gtx, func(gtx C) D {
										return widget.Image{
											Src: paint.NewImageOp(c.Flag),
											Fit: widget.Contain,
										}.Layout(gtx)
									})
								}))
						}))
				})
			})
		}),
		layout.Expanded(func(gtx C) D {
			return c.ctxArea.Layout(gtx, func(gtx C) D {
				gtx.Constraints.Min = image.Point{}
				return component.Menu(th, &c.menu).Layout(gtx)
			})
		}))
}
