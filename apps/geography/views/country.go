package views

import (
	"gioui-experiment/apps/geography/data"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"strconv"
)

type (
	C = layout.Context
	D = layout.Dimensions

	CountryView struct{}
)

func (cv *CountryView) Layout(gtx C, th *material.Theme, country data.Country) D {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(material.Body2(th, country.Cca2).Layout),
		layout.Rigid(material.Body2(th, country.Cca3).Layout),
		layout.Rigid(material.Body2(th, country.Ccn3).Layout),
		layout.Rigid(material.Body2(th, strconv.FormatFloat(country.Area, 'f', -1, 32)).Layout),
		layout.Rigid(func(gtx C) D {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, func(gtx C) D {
					return widget.Image{
						Src: paint.NewImageOp(country.FlagImg),
						Fit: widget.Contain,
					}.Layout(gtx)
				}))
		}))
}
