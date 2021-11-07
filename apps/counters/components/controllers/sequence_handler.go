package controllers

import (
	"gioui-experiment/apps/counters/components/data"
	color "gioui-experiment/custom_themes/colors"
	"gioui-experiment/custom_widgets"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SequenceHandler struct {
	toWhole   widget.Clickable
	toNatural widget.Clickable
	toPrime   widget.Clickable
	toFib     widget.Clickable
}

func (sh *SequenceHandler) Layout(gtx C, th *material.Theme) D {
	cv := data.CurrVals

	toWholesBtn := layout.Rigid(func(gtx C) D {
		for range sh.toWhole.Clicks() {
			sh.handleSequenceType(cv, data.INTEGERS)
		}
		return custom_widgets.LabeledIconBtn{
			Theme:      th,
			BgColor:    g.Colours[color.DEEP_SKY_BLUE],
			LabelColor: g.Colours[color.BLACK],
			Button:     &sh.toWhole,
			Label:      "Z",
			Icon:       nil,
		}.Layout(gtx)
	})
	toNaturalsBtn := layout.Rigid(func(gtx C) D {
		for range sh.toNatural.Clicks() {
			sh.handleSequenceType(cv, data.NATURALS)
		}
		return custom_widgets.LabeledIconBtn{
			Theme:      th,
			BgColor:    g.Colours[color.DEEP_SKY_BLUE],
			LabelColor: g.Colours[color.BLACK],
			Button:     &sh.toNatural,
			Label:      "N",
			Icon:       nil,
		}.Layout(gtx)
	})
	toPrimesBtn := layout.Rigid(func(gtx C) D {
		for range sh.toPrime.Clicks() {
			sh.handleSequenceType(cv, data.PRIMES)
		}
		return custom_widgets.LabeledIconBtn{
			Theme:      th,
			BgColor:    g.Colours[color.DEEP_SKY_BLUE],
			LabelColor: g.Colours[color.BLACK],
			Button:     &sh.toPrime,
			Label:      "Primes",
			Icon:       nil,
		}.Layout(gtx)
	})
	toFibsBtn := layout.Rigid(func(gtx C) D {
		for range sh.toFib.Clicks() {
			sh.handleSequenceType(cv, data.FIBS)
		}
		return custom_widgets.LabeledIconBtn{
			Theme:      th,
			BgColor:    g.Colours[color.DEEP_SKY_BLUE],
			LabelColor: g.Colours[color.BLACK],
			Button:     &sh.toFib,
			Label:      "Fibs",
			Icon:       nil,
		}.Layout(gtx)
	})

	// lay out, horizontally: Z - space - N - space - Primes - space - Fibs
	return layout.Flex{}.Layout(gtx,
		toWholesBtn, g.SpacerX, toNaturalsBtn, g.SpacerX, toPrimesBtn, g.SpacerX, toFibsBtn,
	)
}

func (sh *SequenceHandler) handleSequenceType(cv *data.Generator, target string) {
	cv.Index = 0
	cv.Step = 1
	cv.Start = 1
	cv.Displayed = 1
	cv.SetActiveSequence(target)
}
