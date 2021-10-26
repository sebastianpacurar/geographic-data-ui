package sections

import (
	"gioui-experiment/custom_widgets"
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Top struct {
	changeToWhole   widget.Clickable
	changeToNatural widget.Clickable
	changeToPrime   widget.Clickable
	changeToFib     widget.Clickable
}

var cv = globals.CounterVals

func (t *Top) Layout(th *material.Theme, gtx C) D {
	if t.changeToNatural.Clicked() {
		cv.CurrVal = "unsigned"
		cv.PEnabled = false
	} else if t.changeToWhole.Clicked() {
		cv.CurrVal = "signed"
		cv.PEnabled = false
	} else if t.changeToPrime.Clicked() {
		cv.CurrVal = "unsigned"
		cv.PEnabled = true
	}

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    globals.Colours["deep-sky-blue"],
				LabelColor: globals.Colours["black"],
				Button:     &t.changeToWhole,
				Label:      "Z",
				Icon:       nil,
			}.Layout(gtx)
		}),

		globals.SpacerX,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    globals.Colours["deep-sky-blue"],
				LabelColor: globals.Colours["black"],
				Button:     &t.changeToNatural,
				Label:      "N",
				Icon:       nil,
			}.Layout(gtx)
		}),

		globals.SpacerX,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    globals.Colours["deep-sky-blue"],
				LabelColor: globals.Colours["black"],
				Button:     &t.changeToPrime,
				Label:      "Primes",
				Icon:       nil,
			}.Layout(gtx)
		}),

		globals.SpacerX,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    globals.Colours["deep-sky-blue"],
				LabelColor: globals.Colours["black"],
				Button:     &t.changeToFib,
				Label:      "Fibo",
				Icon:       nil,
			}.Layout(gtx)
		}),
	)
}
