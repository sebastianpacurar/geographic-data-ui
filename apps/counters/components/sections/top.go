package sections

import (
	"gioui-experiment/apps/counters/components/controllers"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Top struct {
		seq controllers.SequenceHandler
	}
)

func (t *Top) Layout(gtx C, th *material.Theme) D {
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return t.seq.Layout(gtx, th)
		}))
}
