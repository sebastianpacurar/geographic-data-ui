package controllers

import (
	"gioui-experiment/apps/geography/data"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const (
	NAME                  = "Name"
	OFFICIAL_NAME         = "Official Name"
	CAPITAL               = "Capital"
	REGION                = "Region"
	SUBREGION             = "Subregion"
	LANGUAGES             = "Languages"
	CONTINENTS            = "Continents"
	IDD_ROOT              = "IDD Root"
	IDD_SUFFIXES          = "IDD Suffixes"
	TOP_LEVEL_DOMAINS     = "Top Level Domains"
	INDEPENDENT           = "Independent"
	STATUS                = "Status"
	UNITED_NATIONS_MEMBER = "United Nations Member"
	LANDLOCKED            = "Landlocked"
	CCA2                  = "CCA 2"
	CCA3                  = "CCA 3"
	CCN3                  = "CCN 3"
	CIOC                  = "IOC Code"
	FIFA                  = "FIFA Code"
	AREA                  = "Area"
	POPULATION            = "Population"
	LATITUDE              = "Latitude"
	LONGITUDE             = "Longitude"
	START_OF_WEEK         = "Start of Week"
	CAR_SIGNS             = "Car Signs"
	CAR_SIDE              = "Car Side"
)

type (
	C = layout.Context
	D = layout.Dimensions

	DisplayedColumns struct {
		viewed data.Country
		list   widget.List
	}
)

// Layout - Lays out the column checkboxes
func (dc *DisplayedColumns) Layout(gtx C, th *material.Theme) D {
	return material.Body1(th, "Under construction").Layout(gtx)
}
