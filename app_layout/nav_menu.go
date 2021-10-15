package app_layout

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"image"
)

type Menu struct {
	Apps []string
}

func (m Menu) Layout(th *material.Theme, gtx C) D {
	if len(m.Apps) == 0 {
		m.Apps = globals.GetAppsNames()
	}
	return layout.Stack{
		Alignment: layout.NW,
	}.Layout(
		gtx,
		layout.Expanded(func(gtx C) D {
			size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
			bar := globals.ColoredArea(
				gtx,
				gtx.Constraints.Constrain(size),
				globals.Colours["sea-green"],
			)
			return bar
		}),
	)

}
