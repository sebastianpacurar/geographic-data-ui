package controllers

import (
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"image"
)

type StatsData struct{}

func (sd *StatsData) Layout(th *material.Theme, gtx C) D {
	return layout.Stack{
		Alignment: layout.NW,
	}.Layout(
		gtx,
		layout.Expanded(func(gtx C) D {
			size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
			bar := g.ColoredArea(
				gtx,
				gtx.Constraints.Constrain(size),
				g.Colours["aero-blue"],
			)
			return bar
		}),
		layout.Stacked(func(gtx C) D {
			return g.Inset.Layout(gtx, func(gtx C) D {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Flexed(1, func(gtx C) D {
						return material.H5(th, "test 1").Layout(gtx)
					}),
					layout.Flexed(1, func(gtx C) D {
						return material.H5(th, "test 2").Layout(gtx)
					}),
				)
			})
		}),
	)
}
