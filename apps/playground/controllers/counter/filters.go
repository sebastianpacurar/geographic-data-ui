package counter

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	Filters struct {
		ascii  widget.Bool
		binary widget.Bool
		octal  widget.Bool
		hex    widget.Bool
	}
)

func (f *Filters) Layout(gtx C, th *material.Theme) D {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
				layout.Rigid(material.CheckBox(th, &f.ascii, "Ascii").Layout),
				layout.Rigid(material.CheckBox(th, &f.binary, "Binary").Layout),
				layout.Rigid(material.CheckBox(th, &f.octal, "Octal").Layout),
				layout.Rigid(material.CheckBox(th, &f.hex, "Hexa").Layout),
			)
		}))
}
