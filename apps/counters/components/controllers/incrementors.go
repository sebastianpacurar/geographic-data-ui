package controllers

import (
	"gioui-experiment/custom_widgets"
	"gioui-experiment/globals"
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
	plusBtn, minusBtn, resetBtn widget.Clickable
}

func (inc *Incrementor) Layout(th *material.Theme, gtx C) D {
	if cv.CurrVal == "signed" {
		parsedLabel = strconv.FormatInt(cv.CountUnit, 10)
		resetLabel = strconv.FormatInt(cv.ResetVal, 10)
	} else if cv.CurrVal == "unsigned" {
		parsedLabel = strconv.FormatUint(cv.UCountUnit, 10)
		resetLabel = strconv.FormatUint(cv.UResetVal, 10)
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
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

							if cv.PEnabled && cv.PCurrIndex == 0 {
								gtx = gtx.Disabled()
							}

							for range inc.minusBtn.Clicks() {
								if cv.PEnabled {
									cv.UCount = cv.PCache[cv.PCurrIndex]
									cv.PCurrIndex += 1
								} else if cv.CurrVal == "signed" {
									cv.Count -= cv.CountUnit
								} else if cv.CurrVal == "unsigned" {
									cv.UCount -= cv.UCountUnit
								}
							}

							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["red"],
									LabelColor: globals.Colours["white"],
									Button:     &inc.minusBtn,
									Icon:       globals.MinusIcon,
									Label:      parsedLabel,
								}.Layout)

						}),
						globals.SpacerX,

						// Reset Button
						layout.Rigid(func(gtx C) D {
							// if count == reset, disable Reset button
							if inc.isDisabled() {
								gtx = gtx.Disabled()
							}

							for range inc.resetBtn.Clicks() {
								if cv.CurrVal == "unsigned" {
									cv.UCount = cv.UResetVal
								} else {
									cv.Count = cv.ResetVal
								}
							}

							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["blue"],
									LabelColor: globals.Colours["white"],
									Button:     &inc.resetBtn,
									Icon:       globals.RefreshIcon,
									Label:      resetLabel,
								}.Layout)
						}),

						globals.SpacerX,

						// Plus Button
						layout.Rigid(func(gtx C) D {
							for range inc.plusBtn.Clicks() {
								if cv.PEnabled {
									cv.UCount = cv.PCache[cv.PCurrIndex]
									cv.PCurrIndex += 1
								} else if cv.CurrVal == "unsigned" {
									cv.UCount += cv.UCountUnit
								} else {
									cv.Count += cv.CountUnit
								}
							}

							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["green"],
									LabelColor: globals.Colours["black"],
									Button:     &inc.plusBtn,
									Icon:       globals.PlusIcon,
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

func (inc *Incrementor) isDisabled() bool {
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

//TODO: implement for int64 boundaries
func handlePlusBtn() {
	if cv.PEnabled {
		cv.UCount = cv.PCache[cv.PCurrIndex]
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
func handleMinusBtn() {
	if cv.PEnabled {
		cv.PCurrIndex -= 1
		cv.UCount = cv.PCache[cv.PCurrIndex]
	} else if cv.CurrVal == "signed" {
		cv.Count -= cv.CountUnit
	} else if cv.CurrVal == "unsigned" {
		cv.UCount -= cv.UCountUnit
	}
}
