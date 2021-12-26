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
		click widget.Clickable
		name  string
		cca2  string
		cca3  string
		ccn3  string
		area  string
	}
)

// Layout -  TODO: Mockup - in progress
func (t *Table) Layout(gtx C, th *material.Theme) D {

	for i := range data.Data {
		t.rows = append(t.rows, row{
			name: data.Data[i].Name.Common,
			cca2: data.Data[i].Cca2,
			cca3: data.Data[i].Cca3,
			ccn3: data.Data[i].Ccn3,
		})
	}

	t.list.Axis = layout.Vertical
	t.list.Alignment = layout.Middle
	border := widget.Border{
		Color: g.Colours[colors.GREY],
		Width: unit.Px(1),
	}

	return material.List(th, &t.list).Layout(gtx, len(data.Data), func(gtx C, i int) D {
		return material.Clickable(gtx, &t.rows[i].click, func(gtx C) D {

			if t.rows[i].click.Clicked() {
				fmt.Println(fmt.Sprintf("%s is clicked", t.rows[i].name))
			}

			return layout.Flex{}.Layout(gtx,
				layout.Flexed(4, func(gtx C) D {
					return border.Layout(gtx, func(gtx C) D {
						return layout.Inset{
							Top:    unit.Dp(2),
							Bottom: unit.Dp(2),
							Left:   unit.Dp(2),
						}.Layout(gtx, func(gtx C) D {
							return layout.Center.Layout(gtx, func(gtx C) D {
								return material.Body1(th, t.rows[i].name).Layout(gtx)
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
								return material.Body1(th, t.rows[i].cca2).Layout(gtx)
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
								return material.Body1(th, t.rows[i].cca3).Layout(gtx)
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
								return material.Body1(th, t.rows[i].ccn3).Layout(gtx)
							})
						})
					})
				}),
				layout.Flexed(2, func(gtx C) D {
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
	})
}
