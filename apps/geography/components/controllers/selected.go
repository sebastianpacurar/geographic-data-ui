package controllers

import (
	"gioui-experiment/apps/geography/components/countries/data"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
)

type (
	SelectedCountries struct {
		pills []pill
		list  widget.List
		api   data.Countries
		wrap  outlay.GridWrap
		count int
	}

	pill struct {
		name    string
		content data.Country
		click   widget.Clickable
	}
)

func (sc *SelectedCountries) Layout(gtx C, th *material.Theme) D {
	// avoid continuous reiteration
	selectedCount := sc.api.GetSelectedCount()
	if sc.count != selectedCount {
		selected := sc.api.GetSelected()
		sc.pills = make([]pill, selectedCount)
		for i := range sc.pills {
			sc.pills[i].name = selected[i].Name.Common
			sc.pills[i].content = selected[i]
		}
		sc.count = selectedCount
	}

	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return sc.wrap.Layout(gtx, sc.count, func(gtx C, i int) D {
				var area D
				area = layout.Inset{
					Top:    unit.Dp(4),
					Right:  unit.Dp(2),
					Bottom: unit.Dp(4),
					Left:   unit.Dp(2),
				}.Layout(gtx, func(gtx C) D {
					return material.Clickable(gtx, &sc.pills[i].click, func(gtx C) D {
						border := widget.Border{
							Color:        g.Colours[colors.GREY],
							CornerRadius: unit.Dp(4),
							Width:        unit.Dp(2),
						}
						return border.Layout(gtx, func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(5),
								Right:  unit.Dp(5),
								Bottom: unit.Dp(5),
								Left:   unit.Dp(5),
							}.Layout(gtx, func(gtx C) D {
								return material.Body1(th, sc.pills[i].name).Layout(gtx)
							})
						})
					})
				})
				return area
			})
		}))
}
