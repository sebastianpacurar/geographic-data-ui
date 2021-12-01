package countries

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	"gioui-experiment/custom_themes/colors"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type (
	Display struct {
		table
		cards []card
		data.Countries
	}

	table struct {
		rowList, columnList widget.List
		rows                []tableRow
	}

	tableRow struct {
		name  string
		list  widget.List
		cells []tableCell
	}

	tableCell struct {
		name   string
		border widget.Border
		layout func(gtx C) D
	}

	card struct {
		name, capital string
		flag          image.Image
	}
)

func (d *Display) LayCard(gtx C, th *material.Theme, country *card) D {
	size := image.Pt(300, 300)
	gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return widget.Border{
				Color:        g.Colours[colors.GREY],
				CornerRadius: unit.Dp(18),
				Width:        unit.Px(2),
			}.Layout(gtx, func(gtx C) D {
				return g.RColoredArea(gtx,
					size,
					unit.Dp(18),
					g.Colours[colors.WHITE],
				)
			})
		}),
		layout.Stacked(func(gtx C) D {
			gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))
			return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceAround}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{}.Layout(gtx,
						layout.Flexed(1, func(gtx C) D {
							return layout.Center.Layout(gtx, func(gtx C) D {
								return material.Body2(th, country.name).Layout(gtx)
							})
						}),
					)
				}))

			//TODO: use image.Decode directly from the request
			//flag processing
			//layout.Rigid(func(gtx C) D {
			//	imgOp := paint.NewImageOp(country.flag)
			//	return layout.Flex{}.Layout(gtx,
			//		layout.Flexed(1, func(gtx C) D {
			//			return
			//		}))
			//	)
			//}))
		}))

}

func (d *Display) Layout(gtx C, th *material.Theme) D {
	err := d.InitCountries()
	if err != nil {
		return material.H2(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	}

	for i := range data.Data {
		d.cards = append(d.cards, card{
			name: data.Data[i].Name.Common,
			//flag: data.Data[i].Flag.Png
		})
	}

	d.table.rowList.Axis = layout.Vertical
	return material.List(th, &d.table.rowList).Layout(gtx, len(data.Data), func(gtx C, rowIndex int) D {
		return d.LayCard(gtx, th, &d.cards[rowIndex])
	})
}
