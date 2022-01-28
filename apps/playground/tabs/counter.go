package tabs

import (
	"gioui-experiment/apps/playground/data"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"strconv"
)

type (
	C = layout.Context
	D = layout.Dimensions

	CounterTab struct{}
)

func (ct *CounterTab) Layout(gtx C, th *material.Theme) D {
	pgv := data.PgVals

	//TODO find a better way and location to handle the Cache population
	if len(pgv.Cache[data.PRIMES]) == 0 {
		pgv.GenPrimes(data.PLIMIT)
	}
	if len(pgv.Cache[data.FIBS]) == 0 {
		pgv.GenFibs(data.FLIMIT)
	}
	seq := pgv.GetActiveSequence()

	// DISPLAYED NUMBER
	return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx C) D {
		var val string
		switch seq {
		case data.PRIMES:
			val = strconv.FormatUint(pgv.Cache[seq][pgv.Primes.Index], 10)
		case data.FIBS:
			val = strconv.FormatUint(pgv.Cache[seq][pgv.Fibonacci.Index], 10)
		case data.NATURALS:
			val = strconv.FormatUint(pgv.Naturals.Displayed, 10)
		case data.INTEGERS:
			val = strconv.FormatUint(pgv.Integers.Displayed, 10)
		}
		return material.H5(th, val).Layout(gtx)
	})
}
