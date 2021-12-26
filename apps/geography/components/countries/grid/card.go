package grid

import (
	"fmt"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
)

type (
	Card struct {
		Name, Capital string
		Cioc, FlagSrc string
		clicked       bool
		Click         widget.Clickable
		flag          image.Image
	}
)

func (c *Card) LayCard(gtx C, th *material.Theme, card *Card) D {
	size := image.Pt(150, 200)
	gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return widget.Border{
				Color:        g.Colours[colors.GREY],
				CornerRadius: unit.Dp(18),
				Width:        unit.Px(2),
			}.Layout(gtx, func(gtx C) D {
				area := material.Clickable(gtx, &card.Click, func(gtx C) D {

					if card.Click.Clicked() {
						fmt.Println(fmt.Sprintf("%s", card.Name))
					}

					return g.RColoredArea(gtx,
						size,
						unit.Dp(18),
						g.Colours[colors.WHITE],
					)
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
									return material.Body2(th, card.Name).Layout(gtx)
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
									return material.Body2(th, card.Capital).Layout(gtx)
								})
							}))
					}))
			})
		}))

}