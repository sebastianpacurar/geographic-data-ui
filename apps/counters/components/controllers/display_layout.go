package controllers

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	DisplayLayout struct {
		radioBtns widget.Enum
	}

	Unit struct {
	}

	Table struct {
	}

	Card struct {
	}
)

func (dl *DisplayLayout) Layout(gtx C, th *material.Theme) D {
	if len(dl.radioBtns.Value) == 0 {
		dl.radioBtns.Value = "s"
	}
	if dl.radioBtns.Changed() {
		switch dl.radioBtns.Value {
		case "s":
			fmt.Println("single")
		case "g":
			fmt.Println("grid")
		case "t":
			fmt.Println("table")
		}
		op.InvalidateOp{}.Add(gtx.Ops)
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(material.RadioButton(th, &dl.radioBtns, "s", "Single Unit").Layout),
		layout.Rigid(material.RadioButton(th, &dl.radioBtns, "g", "Grid").Layout),
		layout.Rigid(material.RadioButton(th, &dl.radioBtns, "t", "Table").Layout),
	)
}
