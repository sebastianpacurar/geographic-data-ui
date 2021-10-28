package controllers

import (
	"gioui-experiment/apps/counters/components/utils"
	"gioui-experiment/custom_widgets"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SequenceHandler struct {
	changeToWhole   widget.Clickable
	changeToNatural widget.Clickable
	changeToPrime   widget.Clickable
	changeToFib     widget.Clickable
}

func (sh *SequenceHandler) Layout(th *material.Theme, gtx C) D {
	cv := utils.CounterVals
	sh.handleNumType(cv)
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    g.Colours["deep-sky-blue"],
				LabelColor: g.Colours["black"],
				Button:     &sh.changeToWhole,
				Label:      "Z",
				Icon:       nil,
			}.Layout(gtx)
		}),

		g.SpacerX,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    g.Colours["deep-sky-blue"],
				LabelColor: g.Colours["black"],
				Button:     &sh.changeToNatural,
				Label:      "N",
				Icon:       nil,
			}.Layout(gtx)
		}),

		g.SpacerX,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    g.Colours["deep-sky-blue"],
				LabelColor: g.Colours["black"],
				Button:     &sh.changeToPrime,
				Label:      "Primes",
				Icon:       nil,
			}.Layout(gtx)
		}),

		g.SpacerX,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return custom_widgets.LabeledIconBtn{
				Theme:      th,
				BgColor:    g.Colours["deep-sky-blue"],
				LabelColor: g.Colours["black"],
				Button:     &sh.changeToFib,
				Label:      "Fibs",
				Icon:       nil,
			}.Layout(gtx)
		}),
	)
}

func (sh SequenceHandler) handleNumType(cv *utils.CurrentValues) {
	if sh.changeToNatural.Clicked() {
		cv.CurrVal = "unsigned"
		cv.UCount = 0
		for k := range cv.ActiveNumType {
			if k == utils.NATURALS {
				cv.ActiveNumType[k] = true
			} else {
				cv.ActiveNumType[k] = false
			}
		}
	} else if sh.changeToWhole.Clicked() {
		cv.CurrVal = "signed"
		cv.Count = 0
		for k := range cv.ActiveNumType {
			if k == utils.WHOLES {
				cv.ActiveNumType[k] = true
			} else {
				cv.ActiveNumType[k] = false
			}
		}
	} else if sh.changeToPrime.Clicked() {
		cv.CurrVal = "unsigned"
		cv.UCount = cv.PCache[0]
		for k := range cv.ActiveNumType {
			if k == utils.PRIMES {
				cv.ActiveNumType[k] = true
			} else {
				cv.ActiveNumType[k] = false
			}
		}
	} else if sh.changeToFib.Clicked() {
		cv.CurrVal = "unsigned"
		cv.UCount = cv.FCache[0]
		for k := range cv.ActiveNumType {
			if k == utils.FIBS {
				cv.ActiveNumType[k] = true
			} else {
				cv.ActiveNumType[k] = false
			}
		}
	}
}
