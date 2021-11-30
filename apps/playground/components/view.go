package components

import (
	"gioui-experiment/apps/playground/components/data"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"strconv"
)

type View struct{}

func (v *View) Layout(th *material.Theme) layout.FlexChild {
	cv := data.CurrVals

	//TODO find a better way and location to handle the Cache population
	if len(cv.Cache[data.PRIMES]) == 0 {
		cv.GenPrimes(data.PLIMIT)
	}
	if len(cv.Cache[data.FIBS]) == 0 {
		cv.GenFibs(data.FLIMIT)
	}
	seq := cv.GetActiveSequence()

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
			case data.PRIMES:
				val = strconv.FormatUint(cv.Cache[seq][cv.Primes.Index], 10)
			case data.FIBS:
				val = strconv.FormatUint(cv.Cache[seq][cv.Fibonacci.Index], 10)
			case data.NATURALS:
				val = strconv.FormatUint(cv.Naturals.Displayed, 10)
			case data.INTEGERS:
				val = strconv.FormatUint(cv.Integers.Displayed, 10)
			}
			return material.H5(th, val).Layout(gtx)
		})
	})
}
