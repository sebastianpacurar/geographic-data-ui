package components

import (
	"gioui-experiment/apps/geography/components/countries"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Geography struct {
	countries.Display
}

func (g *Geography) Layout(th *material.Theme) layout.FlexChild {
	return layout.Rigid(func(gtx C) D {
		return g.Display.Layout(gtx, th)
	})
}
