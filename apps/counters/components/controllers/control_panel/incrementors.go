package control_panel

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

func (inc *Incrementor) Layout(gtx C, th *material.Theme) D {
	cv := data.CurrVals

	minusBtn := layout.Rigid(func(gtx C) D {
		if inc.isMinusBtnDisabled(cv) {
			gtx = gtx.Disabled()
		}
		for range inc.minusBtn.Clicks() {
			inc.handleMinusBtn(cv)
		}
		return g.Inset.Layout(gtx,
			material.IconButton(th, &inc.minusBtn, g.MinusIcon).Layout)
	})

	resetBtn := layout.Rigid(func(gtx C) D {
		if inc.isResetBtnDisabled(cv) {
			gtx = gtx.Disabled()
		}
		for range inc.resetBtn.Clicks() {
			inc.handleResetBtn(cv)
		}
		return g.Inset.Layout(gtx,
			material.IconButton(th, &inc.resetBtn, g.RefreshIcon).Layout)
	})

	plusBtn := layout.Rigid(func(gtx C) D {
		if inc.isPlusBtnDisabled(cv) {
			gtx = gtx.Disabled()
		}
		for range inc.plusBtn.Clicks() {
			inc.handlePlusBtn(cv)
		}
		return g.Inset.Layout(gtx,
			material.IconButton(th, &inc.plusBtn, g.PlusIcon).Layout)
	})

	// laying out minusBtn - space - resetBtn - space - plusBtn
	// align them from the start
	return layout.Flex{Spacing: layout.SpaceEnd}.Layout(gtx,
		minusBtn, g.SpacerX, resetBtn, g.SpacerX, plusBtn,
	)
}

func (inc *Incrementor) isResetBtnDisabled(cv *data.Generator) bool {
	seq := cv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES, data.FIBS:
		res = cv.Cache[seq][cv.Index] == cv.Cache[seq][int(cv.Start)-1]
	case data.NATURALS, data.INTEGERS:
		res = cv.Displayed == cv.Start
	}
	return res
}

func (inc *Incrementor) isMinusBtnDisabled(cv *data.Generator) bool {
	seq := cv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES, data.FIBS:
		res = cv.Index <= 0
	case data.NATURALS, data.INTEGERS:
		res = cv.Displayed <= 0
	}
	return res
}

func (inc *Incrementor) isPlusBtnDisabled(cv *data.Generator) bool {
	seq := cv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES, data.FIBS:
		res = cv.Index == len(cv.Cache[seq])-1
	case data.NATURALS, data.INTEGERS:
		res = cv.Displayed == math.MaxUint64
	}
	return res
}

func (inc *Incrementor) handleResetBtn(cv *data.Generator) {
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		cv.Index = int(cv.Start) - 1
	case data.NATURALS, data.INTEGERS:
		cv.Displayed = cv.Start
	}
}

func (inc *Incrementor) handlePlusBtn(cv *data.Generator) {
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		cv.Index += int(cv.Step)
	case data.NATURALS, data.INTEGERS:
		cv.Displayed += cv.Step
	}
}

func (inc *Incrementor) handleMinusBtn(cv *data.Generator) {
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		cv.Index -= int(cv.Step)
	case data.NATURALS, data.INTEGERS:
		cv.Displayed -= cv.Step
	}
}

func (inc *Incrementor) parseResetLabel(cv *data.Generator) string {
	var lbl string
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		lbl = strconv.FormatUint(cv.Cache[seq][cv.Start-1], 10)
	case data.NATURALS, data.INTEGERS:
		lbl = strconv.FormatUint(cv.Start, 10)
	}
	return lbl
}
