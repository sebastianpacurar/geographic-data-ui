package controllers

import (
	"gioui-experiment/apps/geography/components/countries/data"
	"gioui-experiment/apps/geography/components/countries/grid"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	CountryDetails struct {
		cards []grid.Card
		list  widget.List
	}
)

func (cd *CountryDetails) Layout(gtx C, th *material.Theme) D {
	var content D

	for i := range data.Data {
		if data.Data[i].Hovered {
			content = layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{
						Bottom: unit.Dp(5),
					}.Layout(gtx, func(gtx C) D {
						return material.Body1(th, "test").Layout(gtx)
					})
				}))
		}
	}
	return content
}
