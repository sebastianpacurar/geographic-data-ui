package controllers

import (
	"gioui-experiment/apps/geography/data"
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
		name     string
		content  data.Country
		click    widget.Clickable
		deselect widget.Clickable
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

				if sc.pills[i].deselect.Clicked() {
					name := sc.pills[i].name
					for i := range data.Data {
						if data.Data[i].Name.Common == name {
							data.Data[i].Selected = false
						}
					}
				}

				if sc.pills[i].click.Clicked() {
					name := sc.pills[i].name
					for i := range data.Data {
						if data.Data[i].Name.Common == name {
							data.Data[i].IsCPViewed = true
						} else {
							data.Data[i].IsCPViewed = false
						}
					}
				}

				var area D
				area = layout.Inset{
					Top:    unit.Dp(4),
					Right:  unit.Dp(2),
					Bottom: unit.Dp(4),
					Left:   unit.Dp(2),
				}.Layout(gtx, func(gtx C) D {
					border := widget.Border{
						Color:        g.Colours[colors.GREY],
						CornerRadius: unit.Dp(1),
						Width:        unit.Px(1),
					}
					return border.Layout(gtx, func(gtx C) D {
						return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								var btn material.ButtonStyle
								btn = material.Button(th, &sc.pills[i].click, sc.pills[i].content.Name.Common)
								btn.CornerRadius = unit.Dp(10)
								btn.Background = g.Colours[colors.AERO_BLUE]
								btn.Color = g.Colours[colors.BLACK]
								btn.TextSize = th.TextSize.Scale(14.0 / 16.0)
								return btn.Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								var btn material.ButtonStyle
								btn = material.Button(th, &sc.pills[i].deselect, "X")
								btn.CornerRadius = unit.Dp(0)
								btn.Background = g.Colours[colors.FLAME_RED]
								btn.Color = g.Colours[colors.WHITE]
								btn.TextSize = th.TextSize.Scale(14.0 / 16.0)
								return btn.Layout(gtx)
							}))
					})
				})
				return area
			})
		}))
}
