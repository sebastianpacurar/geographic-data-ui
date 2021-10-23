package counters

import (
	"gioui-experiment/apps/counters/components/sections"
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Page struct {
	Top    sections.Top
	View   sections.View
	Bottom sections.Bottom
}

func (p *Page) Layout(th *material.Theme, gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		layout.Rigid(func(gt C) D {
			return globals.Inset.Layout(gtx, func(gtx C) D {
				return p.Top.Layout(th, gtx)
			})
		}),
		layout.Flexed(1, func(gtx C) D {
			return p.View.Layout(th, gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return globals.Inset.Layout(gtx, func(gtx C) D {
				return p.Bottom.Layout(th, gtx)
			})
		}),
	)
}
