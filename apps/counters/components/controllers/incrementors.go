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

type Incrementor struct {
	plusBtn, minusBtn, resetBtn widget.Clickable
}

func (c *Incrementor) Layout(th *material.Theme, gtx C) D {
	currVal := globals.CurrentNum
	var parsedLabel, resetLabel string
	if currVal == "signed" {
		parsedLabel = strconv.FormatInt(globals.CountUnit, 10)
		resetLabel = strconv.FormatInt(globals.ResetVal, 10)
	} else if currVal == "unsigned" {
		parsedLabel = strconv.FormatUint(globals.UCountUnit, 10)
		resetLabel = strconv.FormatUint(globals.UResetVal, 10)
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
							for range c.minusBtn.Clicks() {
								if currVal == "signed" {
									globals.Count -= globals.CountUnit
								} else if currVal == "unsigned" {
									globals.UCount -= globals.UCountUnit
								}
							}

							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["red"],
									LabelColor: globals.Colours["white"],
									Button:     &c.minusBtn,
									Icon:       globals.MinusIcon,
									Label:      parsedLabel,
								}.Layout)

						}),
						globals.SpacerX,

						// Reset Button
						layout.Rigid(func(gtx C) D {
							// if count == reset, disable Reset button
							if isDisabled() {
								gtx = gtx.Disabled()
							}

							for range c.resetBtn.Clicks() {
								if currVal == "unsigned" {
									globals.UCount = globals.UResetVal
								} else if currVal == "signed" {
									globals.Count = globals.ResetVal
								}
							}

							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["blue"],
									LabelColor: globals.Colours["white"],
									Button:     &c.resetBtn,
									Icon:       globals.RefreshIcon,
									Label:      resetLabel,
								}.Layout)
						}),

						globals.SpacerX,

						// Plus Button
						layout.Rigid(func(gtx C) D {
							for range c.plusBtn.Clicks() {
								if currVal == "unsigned" {
									globals.UCount += globals.UCountUnit
								} else if currVal == "signed" {
									globals.Count += globals.CountUnit
								}
							}

							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["green"],
									LabelColor: globals.Colours["black"],
									Button:     &c.plusBtn,
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

func isDisabled() bool {
	res := true
	switch t := globals.CurrentNum; t {
	case "signed":
		if globals.Count != globals.ResetVal {
			res = false
		}
	case "unsigned":
		if globals.UCount != globals.UResetVal {
			res = false
		}
	}
	return res
}

//TODO: implement for int64 boundaries
func handlePlusBtn() {

}

//TODO: implement mostly for negative values in case of uint64
func handleMinusBtn() {

}
