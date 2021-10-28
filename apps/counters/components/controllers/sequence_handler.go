package controllers

import (
	"gioui-experiment/apps/counters/components/data"
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
	cv := data.CounterVals
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			for range sh.changeToWhole.Clicks() {
				sh.handleSequenceType(cv, data.WHOLES)
			}
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
			for range sh.changeToNatural.Clicks() {
				sh.handleSequenceType(cv, data.NATURALS)
			}
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
			for range sh.changeToPrime.Clicks() {
				sh.handleSequenceType(cv, data.PRIMES)
			}
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
			for range sh.changeToFib.Clicks() {
				sh.handleSequenceType(cv, data.FIBS)
			}
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

func (sh SequenceHandler) handleSequenceType(cv *data.CurrentValues, target string) {
	switch target {
	case data.PRIMES:
		cv.PCurrIndex = 0
		cv.PCount = cv.PCache[cv.PCurrIndex]
		cv.PCountUnit = 1
		cv.PResetVal = 0
	case data.FIBS:
		cv.FCurrIndex = 0
		cv.FCount = cv.FCache[cv.FCurrIndex]
		cv.PCountUnit = 1
		cv.FResetVal = 0
	case data.NATURALS:
		cv.NCount = 0
		cv.NCountUnit = 1
		cv.FResetVal = 0
	case data.WHOLES:
		cv.WCount = 0
		cv.WCountUnit = 1
		cv.FResetVal = 0
	}
	cv.SetActiveSequence(target)
}
