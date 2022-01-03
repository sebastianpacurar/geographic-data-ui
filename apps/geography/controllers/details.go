package controllers

import (
	"fmt"
	"gioui-experiment/apps/geography/data"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions

	CountryDetails struct {
		viewed data.Country
		list   widget.List
	}
)

func (cd *CountryDetails) Layout(gtx C, th *material.Theme) D {
	var content D

	for i := range data.Data {
		if data.Data[i].IsCPViewed {
			cd.viewed = data.Data[i]
		}
	}
	if cd.viewed.Name.Common != "" {
		content = layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return layout.Inset{
					Bottom: unit.Dp(5),
				}.Layout(gtx, func(gtx C) D {
					return cd.LayData(gtx, th, cd.viewed)
				})
			}))
	}

	return content
}

// LayData - Lays all details about the hovered country in the CP
func (cd *CountryDetails) LayData(gtx C, th *material.Theme, country data.Country) D {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(component.DividerSubheadingText(th, "Name").Layout),
		layout.Rigid(material.Body1(th, country.Name.Common).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "Official Name").Layout),
		layout.Rigid(material.Body1(th, country.Name.Official).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "Code ISO 3166-1 alpha-2 (cca2)").Layout),
		layout.Rigid(material.Body1(th, country.Cca2).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "Code ISO 3166-1 alpha-3 (cca3)").Layout),
		layout.Rigid(material.Body1(th, country.Cca3).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "Code ISO 3166-1 numeric (ccn3)").Layout),
		layout.Rigid(material.Body1(th, country.Ccn3).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "International Olympic Committee Code (IOC)").Layout),
		layout.Rigid(material.Body1(th, country.Cioc).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "Area").Layout),
		layout.Rigid(material.Body1(th, fmt.Sprintf("%.2f", country.Area)).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "Population").Layout),
		layout.Rigid(material.Body1(th, fmt.Sprintf("%d", country.Population)).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "Independent").Layout),
		layout.Rigid(func(gtx C) D {
			var res string
			if country.Independent {
				res = "Yes"
			} else {
				res = "No"
			}
			return material.Body1(th, res).Layout(gtx)
		}),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "Independency obtained").Layout),
		layout.Rigid(material.Body1(th, fmt.Sprintf(country.Status)).Layout),
		layout.Rigid(component.Divider(th).Layout),

		layout.Rigid(component.DividerSubheadingText(th, "United Nations Member").Layout),
		layout.Rigid(func(gtx C) D {
			var res string
			if country.UNMember {
				res = "Yes"
			} else {
				res = "No"
			}
			return material.Body1(th, res).Layout(gtx)
		}),
		layout.Rigid(component.Divider(th).Layout),
	)
}
