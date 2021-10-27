package controllers

import (
	"gioui-experiment/apps/counters/components/utils"
	"gioui-experiment/custom_widgets"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
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
	cv := utils.CounterVals
	if cv.CurrVal == "signed" {
		parsedLabel = strconv.FormatInt(cv.CountUnit, 10)
		resetLabel = strconv.FormatInt(cv.ResetVal, 10)
	} else if cv.CurrVal == "unsigned" {
		parsedLabel = strconv.FormatUint(cv.UCountUnit, 10)
		resetLabel = strconv.FormatUint(cv.UResetVal, 10)
	}

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
							// if count == reset, disable Reset button
							if inc.isResetBtnDisabled(cv) {
								gtx = gtx.Disabled()
							}

							for range inc.resetBtn.Clicks() {
								if cv.CurrVal == "unsigned" {
									cv.UCount = cv.UResetVal
								} else {
									cv.Count = cv.ResetVal
								}
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

func (inc Incrementor) isResetBtnDisabled(cv *utils.CurrentValues) bool {
	res := true
	switch cv.CurrVal {
	case "signed":
		if cv.Count != cv.ResetVal {
			res = false
		}
	case "unsigned":
		if cv.UCount != cv.UResetVal {
			res = false
		}
	}
	return res
}

func (inc Incrementor) isMinusBtnDisabled(cv *utils.CurrentValues) bool {
	res := false
	if cv.PEnabled {
		diff := cv.UCount - cv.UCountUnit
		if diff < 0 || cv.PCurrIndex == 0 {
			res = true
		}
	} else if cv.NEnabled {
		diff := cv.UCount - cv.UCountUnit
		if diff < 0 || cv.Count == 0 {
			res = true
		}
	}
	return res
}

//TODO: implement for int64 boundaries
func (inc Incrementor) handlePlusBtn(cv *utils.CurrentValues) {
	if cv.PEnabled {
		cv.UCount = cv.PCache[cv.PCurrIndex+1]
		cv.PCurrIndex += 1
	} else if cv.FEnabled {
		//TODO: to be completed for both buttons
	} else if cv.CurrVal == "signed" {
		cv.Count += cv.CountUnit
	} else if cv.CurrVal == "unsigned" {
		cv.UCount += cv.UCountUnit
	}
}

//TODO: implement mostly for negative values in case of uint64
func (inc Incrementor) handleMinusBtn(cv *utils.CurrentValues) {
	if cv.PEnabled {
		cv.UCount = cv.PCache[cv.PCurrIndex-1]
		cv.PCurrIndex -= 1
	} else if cv.FEnabled {
		//TODO: to be completed for both buttons
	} else if cv.NEnabled || cv.WEnabled {
		cv.Count -= cv.CountUnit
	}
}
