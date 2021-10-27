package sections

import (
	"gioui-experiment/apps/counters/components/utils"
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

func (t *Top) Layout(th *material.Theme, gtx C) D {
	cv := utils.CounterVals
	if t.changeToNatural.Clicked() {
		cv.CurrVal = "unsigned"
		cv.PEnabled = false
		cv.NEnabled = true
		cv.FEnabled = false
		cv.WEnabled = false
	} else if t.changeToWhole.Clicked() {
		cv.CurrVal = "signed"
		cv.PEnabled = false
		cv.NEnabled = false
		cv.FEnabled = false
		cv.WEnabled = true
	} else if t.changeToPrime.Clicked() {
		cv.CurrVal = "unsigned"
		cv.PEnabled = true
		cv.NEnabled = false
		cv.FEnabled = false
		cv.WEnabled = false
	} else if t.changeToFib.Clicked() {
		cv.CurrVal = "unsigned"
		cv.PEnabled = false
		cv.NEnabled = false
		cv.FEnabled = true
		cv.WEnabled = false
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
