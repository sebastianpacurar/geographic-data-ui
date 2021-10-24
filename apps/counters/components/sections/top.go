package sections

import (
	"gioui-experiment/custom_widgets"
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Top struct {
	wholeNums   widget.Clickable
	naturalNums widget.Clickable
	primeNums   widget.Clickable
}

func (t *Top) Layout(th *material.Theme, gtx C) D {
	if t.naturalNums.Clicked() || t.primeNums.Clicked() {
		globals.CurrentNum = "unsigned"
		globals.UCount = 0
		globals.UCountUnit = 1
		globals.UResetVal = 0
	} else if t.wholeNums.Clicked() {
		globals.CurrentNum = "signed"
		globals.Count = 0
		globals.CountUnit = 1
		globals.ResetVal = 0
	}

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    globals.Colours["deep-sky-blue"],
				LabelColor: globals.Colours["black"],
				Button:     &t.wholeNums,
				Label:      "Whole Numbers",
				Icon:       nil,
			}.Layout(gtx)
		}),

		globals.SpacerX,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    globals.Colours["deep-sky-blue"],
				LabelColor: globals.Colours["black"],
				Button:     &t.naturalNums,
				Label:      "Natural Numbers",
				Icon:       nil,
			}.Layout(gtx)
		}),

		globals.SpacerX,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    globals.Colours["deep-sky-blue"],
				LabelColor: globals.Colours["black"],
				Button:     &t.primeNums,
				Label:      "Prime Numbers",
				Icon:       nil,
			}.Layout(gtx)
		}),
	)
}
