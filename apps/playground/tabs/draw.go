package tabs

import (
	"gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"image"
)

type (
	DrawTab struct{}
)

func (dt *DrawTab) Layout(gtx C, th *material.Theme) D {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
			return globals.RColoredArea(gtx, size, 10, globals.Colours[colours.WHITE])
		}),
		layout.Stacked(func(gtx C) D {
			return D{}
		}))
}
