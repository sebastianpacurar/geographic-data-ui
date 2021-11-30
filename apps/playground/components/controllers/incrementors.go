package controllers

import (
	"gioui-experiment/apps/playground/components/data/counter"
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
	pgv := counter.PgVals

	minusBtn := layout.Rigid(func(gtx C) D {
		if inc.isMinusBtnDisabled(pgv) {
			gtx = gtx.Disabled()
		}
		for range inc.minusBtn.Clicks() {
			inc.handleMinusBtn(pgv)
		}
		return g.Inset.Layout(gtx,
			material.IconButton(th, &inc.minusBtn, g.MinusIcon).Layout)
	})

	resetBtn := layout.Rigid(func(gtx C) D {
		if inc.isResetBtnDisabled(pgv) {
			gtx = gtx.Disabled()
		}
		for range inc.resetBtn.Clicks() {
			inc.handleResetBtn(pgv)
		}
		return g.Inset.Layout(gtx,
			material.IconButton(th, &inc.resetBtn, g.RefreshIcon).Layout)
	})

	plusBtn := layout.Rigid(func(gtx C) D {
		if inc.isPlusBtnDisabled(pgv) {
			gtx = gtx.Disabled()
		}
		for range inc.plusBtn.Clicks() {
			inc.handlePlusBtn(pgv)
		}
		return g.Inset.Layout(gtx,
			material.IconButton(th, &inc.plusBtn, g.PlusIcon).Layout)
	})

	// laying out minusBtn - space - resetBtn - space - plusBtn
	// align them from the start
	return layout.Flex{Spacing: layout.SpaceEvenly}.Layout(gtx,
		minusBtn, g.SpacerX, resetBtn, g.SpacerX, plusBtn,
	)
}

func (inc *Incrementor) isResetBtnDisabled(pgv *counter.Generator) bool {
	seq := pgv.GetActiveSequence()
	res := false
	switch seq {
	case counter.PRIMES:
		res = pgv.Cache[seq][pgv.Primes.Index] == pgv.Cache[seq][int(pgv.Primes.Start)-1]
	case counter.FIBS:
		res = pgv.Cache[seq][pgv.Fibonacci.Index] == pgv.Cache[seq][int(pgv.Fibonacci.Start)-1]
	case counter.NATURALS:
		res = pgv.Naturals.Displayed == pgv.Naturals.Start
	case counter.INTEGERS:
		res = pgv.Integers.Displayed == pgv.Integers.Start
	}
	return res
}

func (inc *Incrementor) isMinusBtnDisabled(pgv *counter.Generator) bool {
	seq := pgv.GetActiveSequence()
	res := false
	switch seq {
	case counter.PRIMES:
		res = pgv.Primes.Index <= 0
	case counter.FIBS:
		res = pgv.Fibonacci.Index <= 0
	case counter.NATURALS:
		res = pgv.Naturals.Displayed <= 0
	case counter.INTEGERS:
		res = pgv.Integers.Displayed <= 0
	}
	return res
}

func (inc *Incrementor) isPlusBtnDisabled(pgv *counter.Generator) bool {
	seq := pgv.GetActiveSequence()
	res := false
	switch seq {
	case counter.PRIMES:
		res = pgv.Primes.Index == len(pgv.Cache[seq])-1
	case counter.FIBS:
		res = pgv.Fibonacci.Index == len(pgv.Cache[seq])-1
	case counter.NATURALS:
		res = pgv.Naturals.Displayed == math.MaxUint64
	case counter.INTEGERS:
		res = pgv.Integers.Displayed == math.MaxUint64
	}
	return res
}

func (inc *Incrementor) handleResetBtn(pgv *counter.Generator) {
	seq := pgv.GetActiveSequence()
	switch seq {
	case counter.PRIMES:
		pgv.Primes.Index = int(pgv.Primes.Start) - 1
	case counter.FIBS:
		pgv.Fibonacci.Index = int(pgv.Fibonacci.Start) - 1
	case counter.NATURALS:
		pgv.Naturals.Displayed = pgv.Naturals.Start
	case counter.INTEGERS:
		pgv.Integers.Displayed = pgv.Integers.Start
	}
}

func (inc *Incrementor) handlePlusBtn(pgv *counter.Generator) {
	seq := pgv.GetActiveSequence()
	switch seq {
	case counter.PRIMES:
		pgv.Primes.Index += int(pgv.Primes.Step)
	case counter.FIBS:
		pgv.Fibonacci.Index += int(pgv.Fibonacci.Step)
	case counter.NATURALS:
		pgv.Naturals.Displayed += pgv.Naturals.Step
	case counter.INTEGERS:
		pgv.Integers.Displayed += pgv.Integers.Step
	}
}

func (inc *Incrementor) handleMinusBtn(pgv *counter.Generator) {
	seq := pgv.GetActiveSequence()
	switch seq {
	case counter.PRIMES:
		pgv.Primes.Index -= int(pgv.Primes.Step)
	case counter.FIBS:
		pgv.Fibonacci.Index -= int(pgv.Fibonacci.Step)
	case counter.NATURALS:
		pgv.Naturals.Displayed -= pgv.Naturals.Step
	case counter.INTEGERS:
		pgv.Integers.Displayed -= pgv.Integers.Step
	}
}

func (inc *Incrementor) parseResetLabel(pgv *counter.Generator) string {
	var lbl string
	seq := pgv.GetActiveSequence()
	switch seq {
	case counter.PRIMES:
		lbl = strconv.FormatUint(pgv.Cache[seq][pgv.Primes.Start-1], 10)
	case counter.FIBS:
		lbl = strconv.FormatUint(pgv.Cache[seq][pgv.Fibonacci.Start-1], 10)
	case counter.NATURALS:
		lbl = strconv.FormatUint(pgv.Naturals.Start, 10)
	case counter.INTEGERS:
		lbl = strconv.FormatUint(pgv.Integers.Start, 10)
	}
	return lbl
}
