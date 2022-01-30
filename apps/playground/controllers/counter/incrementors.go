package counter

import (
	"gioui-experiment/apps/playground/data"
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
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
	pgv := data.PgVals

	minusBtn := layout.Rigid(func(gtx C) D {
		if inc.isMinusBtnDisabled(pgv) {
			gtx = gtx.Disabled()
		}
		for range inc.minusBtn.Clicks() {
			inc.handleMinusBtn(pgv)
		}
		return layout.UniformInset(unit.Dp(10)).Layout(gtx,
			material.IconButton(th, &inc.minusBtn, globals.MinusIcon, "desc").Layout)
	})

	resetBtn := layout.Rigid(func(gtx C) D {
		if inc.isResetBtnDisabled(pgv) {
			gtx = gtx.Disabled()
		}
		for range inc.resetBtn.Clicks() {
			inc.handleResetBtn(pgv)
		}
		return layout.UniformInset(unit.Dp(10)).Layout(gtx,
			material.IconButton(th, &inc.resetBtn, globals.RefreshIcon, "desc").Layout)
	})

	plusBtn := layout.Rigid(func(gtx C) D {
		if inc.isPlusBtnDisabled(pgv) {
			gtx = gtx.Disabled()
		}
		for range inc.plusBtn.Clicks() {
			inc.handlePlusBtn(pgv)
		}
		return layout.UniformInset(unit.Dp(10)).Layout(gtx,
			material.IconButton(th, &inc.plusBtn, globals.PlusIcon, "desc").Layout)
	})

	// laying out minusBtn - space - resetBtn - space - plusBtn
	// align them from the start
	return layout.Flex{Spacing: layout.SpaceEvenly}.Layout(gtx,
		minusBtn,
		layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
		resetBtn,
		layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
		plusBtn,
	)
}

func (inc *Incrementor) isResetBtnDisabled(pgv *data.Generator) bool {
	seq := pgv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES:
		res = pgv.Cache[seq][pgv.Primes.Index] == pgv.Cache[seq][int(pgv.Primes.Start)-1]
	case data.FIBS:
		res = pgv.Cache[seq][pgv.Fibonacci.Index] == pgv.Cache[seq][int(pgv.Fibonacci.Start)-1]
	case data.NATURALS:
		res = pgv.Naturals.Displayed == pgv.Naturals.Start
	case data.INTEGERS:
		res = pgv.Integers.Displayed == pgv.Integers.Start
	}
	return res
}

func (inc *Incrementor) isMinusBtnDisabled(pgv *data.Generator) bool {
	seq := pgv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES:
		res = pgv.Primes.Index <= 0
	case data.FIBS:
		res = pgv.Fibonacci.Index <= 0
	case data.NATURALS:
		res = pgv.Naturals.Displayed <= 0
	case data.INTEGERS:
		res = pgv.Integers.Displayed <= 0
	}
	return res
}

func (inc *Incrementor) isPlusBtnDisabled(pgv *data.Generator) bool {
	seq := pgv.GetActiveSequence()
	res := false
	switch seq {
	case data.PRIMES:
		res = pgv.Primes.Index == len(pgv.Cache[seq])-1
	case data.FIBS:
		res = pgv.Fibonacci.Index == len(pgv.Cache[seq])-1
	case data.NATURALS:
		res = pgv.Naturals.Displayed == math.MaxUint64
	case data.INTEGERS:
		res = pgv.Integers.Displayed == math.MaxUint64
	}
	return res
}

func (inc *Incrementor) handleResetBtn(pgv *data.Generator) {
	seq := pgv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		pgv.Primes.Index = int(pgv.Primes.Start) - 1
	case data.FIBS:
		pgv.Fibonacci.Index = int(pgv.Fibonacci.Start) - 1
	case data.NATURALS:
		pgv.Naturals.Displayed = pgv.Naturals.Start
	case data.INTEGERS:
		pgv.Integers.Displayed = pgv.Integers.Start
	}
}

func (inc *Incrementor) handlePlusBtn(pgv *data.Generator) {
	seq := pgv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		pgv.Primes.Index += int(pgv.Primes.Step)
	case data.FIBS:
		pgv.Fibonacci.Index += int(pgv.Fibonacci.Step)
	case data.NATURALS:
		pgv.Naturals.Displayed += pgv.Naturals.Step
	case data.INTEGERS:
		pgv.Integers.Displayed += pgv.Integers.Step
	}
}

func (inc *Incrementor) handleMinusBtn(pgv *data.Generator) {
	seq := pgv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		pgv.Primes.Index -= int(pgv.Primes.Step)
	case data.FIBS:
		pgv.Fibonacci.Index -= int(pgv.Fibonacci.Step)
	case data.NATURALS:
		pgv.Naturals.Displayed -= pgv.Naturals.Step
	case data.INTEGERS:
		pgv.Integers.Displayed -= pgv.Integers.Step
	}
}

func (inc *Incrementor) parseResetLabel(pgv *data.Generator) string {
	var lbl string
	seq := pgv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		lbl = strconv.FormatUint(pgv.Cache[seq][pgv.Primes.Start-1], 10)
	case data.FIBS:
		lbl = strconv.FormatUint(pgv.Cache[seq][pgv.Fibonacci.Start-1], 10)
	case data.NATURALS:
		lbl = strconv.FormatUint(pgv.Naturals.Start, 10)
	case data.INTEGERS:
		lbl = strconv.FormatUint(pgv.Integers.Start, 10)
	}
	return lbl
}
