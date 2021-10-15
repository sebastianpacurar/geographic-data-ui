package app_layout

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Menu struct{}

func (m Menu) Layout(th *material.Theme, gtx C) D {
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx C) D {
			menu := material.H3(th, "Test")
			return menu.Layout(gtx)
		}),
	)

}
