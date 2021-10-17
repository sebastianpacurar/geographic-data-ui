package counters

import (
	"gioui-experiment/apps/counters/components"
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Page struct {
	TopController    components.ValueHandler
	Viewer           components.View
	BottomController components.Counter
}

func (p *Page) Layout(th *material.Theme, gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx C) D {
			return globals.Inset.Layout(gtx, func(gtx C) D {
				return p.TopController.Layout(th, gtx)
			})
		}),
		layout.Flexed(1, func(gtx C) D {
			return p.Viewer.Layout(th, gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return globals.Inset.Layout(gtx, func(gtx C) D {
				return p.BottomController.Layout(th, gtx)
			})
		}),
	)
}
