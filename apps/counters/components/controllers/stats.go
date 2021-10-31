package controllers

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type StatsData struct{}

func (sd *StatsData) Layout(th *material.Theme, gtx C) D {
	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Baseline,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return material.H5(th, "test 1").Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return material.H5(th, "test 2").Layout(gtx)
		}),
	)
}
