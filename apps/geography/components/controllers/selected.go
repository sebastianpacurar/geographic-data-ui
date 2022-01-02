package controllers

import (
	"gioui-experiment/apps/geography/components/countries/data"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	SelectedCountries struct {
		list widget.List
		tabs []tab
	}

	tab struct {
		selected data.Country
		click    widget.Clickable
	}
)

func (sc *SelectedCountries) Layout(gtx C, th *material.Theme) D {
	sc.list.Alignment = layout.Middle

	for i := range data.Data {
		if data.Data[i].Selected {
			sc.tabs = append(sc.tabs, tab{
				selected: data.Data[i],
			})
		}
	}

	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			border := widget.Border{
				Color: g.Colours[colors.GREY],
				Width: unit.Dp(2),
			}
			return layout.Inset{
				Bottom: unit.Dp(5),
			}.Layout(gtx, func(gtx C) D {
				return material.List(th, &sc.list).Layout(gtx, len(sc.tabs), func(gtx C, i int) D {
					return border.Layout(gtx, func(gtx C) D {
						return layout.Inset{
							Top:    unit.Dp(3),
							Right:  unit.Dp(3),
							Bottom: unit.Dp(3),
							Left:   unit.Dp(3),
						}.Layout(gtx, func(gtx C) D {
							return material.Clickable(gtx, &sc.tabs[i].click, func(gtx C) D {
								return material.Body1(th, sc.tabs[i].selected.Name.Common).Layout(gtx)
							})
						})
					})
				})
			})
		}))
}
