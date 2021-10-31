package controllers

import (
	"gioui-experiment/apps/counters/components/data"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"math"
	"strconv"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Incrementor struct {
		plusBtn  widget.Clickable
		minusBtn widget.Clickable
		resetBtn widget.Clickable
	}
)

func (inc *Incrementor) Layout(th *material.Theme, gtx C) D {
	cv := data.CounterVals
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceEnd,
	}.Layout(gtx,

		// MinusButton
		layout.Rigid(func(gtx C) D {
			if inc.isMinusBtnDisabled(cv) {
				gtx = gtx.Disabled()
			}
			for range inc.minusBtn.Clicks() {
				inc.handleMinusBtn(cv)
			}
			return g.Inset.Layout(gtx,
				material.IconButton(th, &inc.minusBtn, g.MinusIcon).Layout)
		}),
		g.SpacerX,

		// Reset Button
		layout.Rigid(func(gtx C) D {
			if inc.isResetBtnDisabled(cv) {
				gtx = gtx.Disabled()
			}
			for range inc.resetBtn.Clicks() {
				inc.handleResetBtn(cv)
			}
			return g.Inset.Layout(gtx,
				material.IconButton(th, &inc.resetBtn, g.RefreshIcon).Layout)
		}),

		g.SpacerX,

		// Plus Button
		layout.Rigid(func(gtx C) D {
			if inc.isPlusBtnDisabled(cv) {
				gtx = gtx.Disabled()
			}
			for range inc.plusBtn.Clicks() {
				inc.handlePlusBtn(cv)
			}
			return g.Inset.Layout(gtx,
				material.IconButton(th, &inc.plusBtn, g.PlusIcon).Layout)
		}),
	)
}

func (inc *Incrementor) isResetBtnDisabled(cv *data.CurrentValues) bool {
	seq := cv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES, data.FIBS:
		res = cv.Cache[seq][cv.Index] == cv.Cache[seq][int(cv.Start)-1]
	case data.NATURALS, data.WHOLES:
		res = cv.Displayed == cv.Start
	}
	return res
}

func (inc *Incrementor) isMinusBtnDisabled(cv *data.CurrentValues) bool {
	seq := cv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES, data.FIBS:
		res = cv.Index <= 0
	case data.NATURALS, data.WHOLES:
		res = cv.Displayed <= 0
	}
	return res
}

func (inc *Incrementor) isPlusBtnDisabled(cv *data.CurrentValues) bool {
	seq := cv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES, data.FIBS:
		res = cv.Index == len(cv.Cache[seq])-1
	case data.NATURALS, data.WHOLES:
		res = cv.Displayed == math.MaxUint64
	}
	return res
}

func (inc *Incrementor) handleResetBtn(cv *data.CurrentValues) {
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		cv.Index = int(cv.Start) - 1
	case data.NATURALS, data.WHOLES:
		cv.Displayed = cv.Start
	}
}

func (inc *Incrementor) handlePlusBtn(cv *data.CurrentValues) {
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		cv.Index += int(cv.Step)
	case data.NATURALS, data.WHOLES:
		cv.Displayed += cv.Step
	}
}

func (inc *Incrementor) handleMinusBtn(cv *data.CurrentValues) {
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		cv.Index -= int(cv.Step)
	case data.NATURALS, data.WHOLES:
		cv.Displayed -= cv.Step
	}
}

func (inc *Incrementor) parseResetLabel(cv *data.CurrentValues) string {
	var lbl string
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		lbl = strconv.FormatUint(cv.Cache[seq][cv.Start-1], 10)
	case data.NATURALS, data.WHOLES:
		lbl = strconv.FormatUint(cv.Start, 10)
	}
	return lbl
}
