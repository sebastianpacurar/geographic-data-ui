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

func (t *Top) Layout(th *material.Theme, gtx C) D {
	return t.seq.Layout(th, gtx)
}
