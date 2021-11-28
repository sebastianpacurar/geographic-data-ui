package controllers

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	Filters struct {
		ascii       widget.Bool
		binary      widget.Bool
		octal       widget.Bool
		hexadecimal widget.Bool
	}
)

func (f *Filters) Layout(gtx C, th *material.Theme) D {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return material.CheckBox(th, &f.ascii, "Ascii").Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return material.CheckBox(th, &f.binary, "Binary").Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return material.CheckBox(th, &f.octal, "Octal").Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return material.CheckBox(th, &f.hexadecimal, "Hexadecimal").Layout(gtx)
		}),
	)
}
