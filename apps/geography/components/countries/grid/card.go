package grid

import (
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
)

type (
	Card struct {
		Name     string
		Cca2     string
		Active   bool
		Hovered  bool
		Selected bool
		Click    widget.Clickable
		flag     image.Image

		// menu - triggered by right click
		menu            component.MenuState
		ctxArea         component.ContextArea
		isMenuTriggered bool

		// menu options
		selectBtn     widget.Clickable
		deselectBtn   widget.Clickable
		copyToClipBtn widget.Clickable
	}
)

func (c *Card) LayCard(gtx C, th *material.Theme) D {
	size := image.Pt(200, 275)
	gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))

	if !c.isMenuTriggered {
		lbl := "Select"
		btn := &c.selectBtn
		if c.Selected {
			lbl = "Deselect"
			btn = &c.deselectBtn
		}
		c.menu = component.MenuState{
			Options: []func(gtx C) D{
				component.MenuItem(th, btn, lbl).Layout,
				component.MenuItem(th, &c.copyToClipBtn, "Copy as JSON").Layout,
			},
		}
	}
	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx C) D {
			return widget.Border{
				Color:        g.Colours[colors.GREY],
				CornerRadius: unit.Dp(2),
				Width:        unit.Px(2),
			}.Layout(gtx, func(gtx C) D {

				area := material.Clickable(gtx, &c.Click, func(gtx C) D {
					cardColor := g.Colours[colors.WHITE]

					if c.Selected {
						cardColor = g.Colours[colors.AERO_BLUE]
					}

					if c.Click.Hovered() && !c.Selected {
						cardColor = g.Colours[colors.NYANZA]
					} else if c.Click.Hovered() && c.Selected {
						cardColor = g.Colours[colors.LIGHT_SALMON]
					}

					return g.RColoredArea(gtx, size, 10, cardColor)
				})
				return area
			})
		}),
		layout.Stacked(func(gtx C) D {
			return g.Inset.Layout(gtx, func(gtx C) D {
				gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))
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

					// TODO: fix this or find a workaround
					// (capital area) country flag (temporary broken)
					layout.Rigid(func(gtx C) D {
						//country.flag = d.processFlagFromURL(country)
						return layout.Flex{}.Layout(gtx,
							layout.Flexed(1, func(gtx C) D {
								return layout.Center.Layout(gtx, func(gtx C) D {
									return material.Body2(th, c.Cca2).Layout(gtx)
								})
							}))
					}))
			})
		}),

		// this returns the menu
		layout.Expanded(func(gtx C) D {
			return c.ctxArea.Layout(gtx, func(gtx C) D {
				gtx.Constraints.Min = image.Point{}
				return component.Menu(th, &c.menu).Layout(gtx)
			})
		}),
	)
}
