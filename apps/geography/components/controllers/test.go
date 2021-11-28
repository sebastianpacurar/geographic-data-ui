package controllers

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	TestController struct{}
)

func (tc *TestController) Layout(gtx C, th *material.Theme) D {
	return material.Body2(th, "Test Controller geography").Layout(gtx)
}
