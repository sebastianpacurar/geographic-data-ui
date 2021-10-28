package sections

import (
	"gioui-experiment/apps/counters/components/controllers"
	"gioui.org/widget/material"
)

type Top struct {
	seq controllers.SequenceHandler
}

func (t *Top) Layout(th *material.Theme, gtx C) D {
	return t.seq.Layout(th, gtx)
}
