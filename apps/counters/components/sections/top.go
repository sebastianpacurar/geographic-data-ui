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

	if t.naturalNums.Clicked() {

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
			}.Layout(gtx)
		}),
	)
}
