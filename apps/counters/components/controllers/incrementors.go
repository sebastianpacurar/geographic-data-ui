package controllers

import (
	"gioui-experiment/apps/counters/components/data"
	"gioui-experiment/custom_widgets"
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
)

var (
	parsedLabel string
	resetLabel  string
)

type Incrementor struct {
	plusBtn  widget.Clickable
	minusBtn widget.Clickable
	resetBtn widget.Clickable
}

func (inc *Incrementor) Layout(th *material.Theme, gtx C) D {
	cv := data.CounterVals
	inc.GetLabels(cv)

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceEvenly,
					}.Layout(gtx,
						layout.Rigid(func(gtx C) D {

							if inc.isMinusBtnDisabled(cv) {
								gtx = gtx.Disabled()
							}

							for range inc.minusBtn.Clicks() {
								inc.handleMinusBtn(cv)
							}

							return g.Inset.Layout(gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    g.Colours["red"],
									LabelColor: g.Colours["white"],
									Button:     &inc.minusBtn,
									Icon:       g.MinusIcon,
									Label:      parsedLabel,
								}.Layout)

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
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    g.Colours["blue"],
									LabelColor: g.Colours["white"],
									Button:     &inc.resetBtn,
									Icon:       g.RefreshIcon,
									Label:      resetLabel,
								}.Layout)
						}),

						g.SpacerX,

						// Plus Button
						layout.Rigid(func(gtx C) D {
							for range inc.plusBtn.Clicks() {
								inc.handlePlusBtn(cv)
							}

							return g.Inset.Layout(gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    g.Colours["green"],
									LabelColor: g.Colours["black"],
									Button:     &inc.plusBtn,
									Icon:       g.PlusIcon,
									Label:      parsedLabel,
								}.Layout,
							)
						}),
					)
				}),
			)
		}),
	)
}

func (inc Incrementor) isResetBtnDisabled(cv *data.CurrentValues) bool {
	res := false
	switch cv.GetActiveSequence() {
	case data.PRIMES:
		res = cv.PCount == cv.PCache[cv.PResetVal]
	case data.FIBS:
		res = cv.FCount == cv.FCache[cv.FResetVal]
	case data.NATURALS:
		res = cv.NCount == cv.NResetVal
	case data.WHOLES:
		res = cv.WCount == cv.WResetVal
	}
	return res
}

func (inc Incrementor) isMinusBtnDisabled(cv *data.CurrentValues) bool {
	res := false
	switch cv.GetActiveSequence() {
	case data.PRIMES:
		res = cv.PCurrIndex == 0 || (cv.PCount-cv.PCountUnit) < 0
	case data.FIBS:
		res = cv.FCurrIndex == 0 || (cv.FCount-cv.FCountUnit) < 0
	case data.NATURALS:
		res = cv.NCount == 0 || (cv.NCount-cv.NCountUnit) < 0
	case data.WHOLES:
		// don't let Wholes Counter jump from math.MinInt(64) to math.MaxInt(64) due to
		//  boundary limitations
		res = cv.WCount == math.MinInt64 || (cv.WCount-cv.WCount) > 0
	}
	return res
}

func (inc Incrementor) handleResetBtn(cv *data.CurrentValues) {
	switch cv.GetActiveSequence() {
	case data.PRIMES:
		cv.PCurrIndex = int(cv.PResetVal)
		cv.PCount = cv.PCache[cv.PCurrIndex]
	case data.FIBS:
		cv.FCurrIndex = int(cv.FResetVal)
		cv.FCount = cv.FCache[cv.FCurrIndex]
	case data.NATURALS:
		cv.NCount = cv.NResetVal
	case data.WHOLES:
		cv.WCount = cv.WResetVal
	}
}

func (inc Incrementor) handlePlusBtn(cv *data.CurrentValues) {
	switch cv.GetActiveSequence() {
	case data.PRIMES:
		cv.PCount = cv.PCache[cv.PCurrIndex+1]
		cv.PCurrIndex++
	case data.FIBS:
		cv.FCount = cv.FCache[cv.FCurrIndex+1]
		cv.FCurrIndex++
	case data.NATURALS:
		cv.NCount += cv.NCountUnit
	case data.WHOLES:
		cv.WCount += cv.WCountUnit
	}
}

func (inc Incrementor) handleMinusBtn(cv *data.CurrentValues) {
	switch cv.GetActiveSequence() {
	case data.PRIMES:
		cv.PCount = cv.PCache[cv.PCurrIndex-1]
		cv.PCurrIndex--
	case data.FIBS:
		cv.FCount = cv.FCache[cv.FCurrIndex-1]
		cv.FCurrIndex--
	case data.NATURALS:
		cv.NCount -= cv.NCountUnit
	case data.WHOLES:
		cv.WCount -= cv.WCountUnit
	}
}

func (inc Incrementor) GetLabels(cv *data.CurrentValues) {
	switch cv.GetActiveSequence() {
	case data.PRIMES:
		parsedLabel = strconv.FormatUint(cv.PCountUnit, 10)
		resetLabel = strconv.FormatUint(cv.PResetVal, 10)
	case data.FIBS:
		parsedLabel = strconv.FormatUint(cv.FCountUnit, 10)
		resetLabel = strconv.FormatUint(cv.FResetVal, 10)
	case data.NATURALS:
		parsedLabel = strconv.FormatUint(cv.NCountUnit, 10)
		resetLabel = strconv.FormatUint(cv.NResetVal, 10)
	case data.WHOLES:
		parsedLabel = strconv.FormatInt(cv.WCountUnit, 10)
		resetLabel = strconv.FormatInt(cv.WResetVal, 10)
	}
}
