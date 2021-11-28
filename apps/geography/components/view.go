package components

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Geography struct{}

func (g *Geography) Layout(th *material.Theme) layout.FlexChild {
	return layout.Rigid(func(gtx C) D {
		test := material.H2(th, "Geography Application")
		return test.Layout(gtx)
	})
}
