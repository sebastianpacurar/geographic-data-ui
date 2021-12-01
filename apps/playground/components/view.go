package components

import (
	"gioui-experiment/apps/playground/components/data/counter"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"strconv"
)

type View struct{}

func (v *View) Layout(th *material.Theme) layout.FlexChild {
	pgv := counter.PgVals

	//TODO find a better way and location to handle the Cache population
	if len(pgv.Cache[counter.PRIMES]) == 0 {
		pgv.GenPrimes(counter.PLIMIT)
	}
	if len(pgv.Cache[counter.FIBS]) == 0 {
		pgv.GenFibs(counter.FLIMIT)
	}
	seq := pgv.GetActiveSequence()

	/// DISPLAYED NUMBER
	return layout.Rigid(func(gtx C) D {
		return layout.Inset{
			Top:    unit.Dp(10),
			Right:  unit.Dp(50),
			Bottom: unit.Dp(20),
			Left:   unit.Dp(50),
		}.Layout(gtx, func(gtx C) D {
			var val string
			switch seq {
			case counter.PRIMES:
				val = strconv.FormatUint(pgv.Cache[seq][pgv.Primes.Index], 10)
			case counter.FIBS:
				val = strconv.FormatUint(pgv.Cache[seq][pgv.Fibonacci.Index], 10)
			case counter.NATURALS:
				val = strconv.FormatUint(pgv.Naturals.Displayed, 10)
			case counter.INTEGERS:
				val = strconv.FormatUint(pgv.Integers.Displayed, 10)
			}
			return material.H5(th, val).Layout(gtx)
		})
	})
}
