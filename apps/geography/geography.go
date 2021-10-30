package geography

import (
	"gioui-experiment/apps/geography/components/countries"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Page struct {
		display   countries.Display
		Countries countries.Countries
	}
)

func (p *Page) Layout(th *material.Theme, gtx C) D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			test := material.H2(th, "Geography Page")
			return test.Layout(gtx)
		}),

		g.SpacerX,

		layout.Rigid(func(gtx C) D {
			return p.display.Layout(th, gtx)
		}),
	)
}
