package table

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Table struct {
		rows []row
		list widget.List
	}

	row struct {
		Name    string
		Capital string
		Cioc    string
	}
)

// Layout -  TODO: Mockup - in progress
func (t *Table) Layout(gtx C, th *material.Theme) D {
	t.list.Axis = layout.Vertical
	t.list.Alignment = layout.Middle
	border := widget.Border{
		Color: g.Colours[colors.GREY],
		Width: unit.Px(1),
	}

	return material.List(th, &t.list).Layout(gtx, len(data.Data), func(gtx C, i int) D {
		return layout.Flex{}.Layout(gtx,
			layout.Flexed(1, func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Center.Layout(gtx, func(gtx C) D {
							return material.Body1(th, data.Data[i].Name.Common).Layout(gtx)
						})
					})
				})
			}),
			layout.Flexed(1, func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Center.Layout(gtx, func(gtx C) D {
							return material.Body1(th, data.Data[i].Cca2).Layout(gtx)
						})
					})
				})
			}),
			layout.Flexed(1, func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Center.Layout(gtx, func(gtx C) D {
							return material.Body1(th, data.Data[i].Cca3).Layout(gtx)
						})
					})
				})
			}),
			layout.Flexed(1, func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Center.Layout(gtx, func(gtx C) D {
							return material.Body1(th, data.Data[i].Ccn3).Layout(gtx)
						})
					})
				})
			}),
			layout.Flexed(1, func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Center.Layout(gtx, func(gtx C) D {
							return material.Body1(th, fmt.Sprintf("%.2f", data.Data[i].Area)).Layout(gtx)
						})
					})
				})
			}),
		)
	})
}
