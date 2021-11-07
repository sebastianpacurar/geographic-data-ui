package controllers

import (
	"gioui-experiment/custom_themes/colors"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
)

type ControlPanel struct {
	vh ValueHandler
}

func (cp *ControlPanel) Layout(gtx C, th *material.Theme) D {
	return layout.Stack{Alignment: layout.N}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
			bar := g.ColoredArea(
				gtx,
				gtx.Constraints.Constrain(size),
				g.Colours[colors.AERO_BLUE],
			)
			return bar
		}),
		layout.Stacked(func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,

				// STATS AREA STARTS HERE
				/// Key - Value pair design for Stats
				layout.Rigid(func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(10),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(20),
					}.Layout(gtx, func(gtx C) D {
						return cp.vh.Layout(gtx, th)
					})
				}),
			)
		}))
}
