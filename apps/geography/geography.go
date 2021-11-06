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

	Application struct {
		display   countries.Display
		Countries countries.Countries
	}
)

func (app *Application) Layout(gtx C, th *material.Theme) D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			test := material.H2(th, "Geography Application")
			return test.Layout(gtx)
		}),

		g.SpacerX,

		layout.Rigid(func(gtx C) D {
			return app.display.Layout(gtx, th)
		}),
	)
}
