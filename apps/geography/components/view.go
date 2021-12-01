package components

import (
	"gioui-experiment/apps/geography/components/countries"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Geography struct {
	countries.Display
}

func (geo *Geography) Layout(th *material.Theme) layout.FlexChild {
	return layout.Rigid(func(gtx C) D {
		return g.Inset.Layout(gtx, func(gtx C) D {
			return geo.Display.Layout(gtx, th)
		})
	})
}
