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
								globals.Count -= globals.CountUnit
							}

							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["red"],
									LabelColor: globals.Colours["white"],
									Button:     &c.minusBtn,
									Icon:       globals.MinusIcon,
									Label:      strconv.FormatInt(globals.CountUnit, 10),
								}.Layout)

						}),
						globals.SpacerX,

						// Reset Button
						layout.Rigid(func(gtx C) D {
							// if count == reset, disable Reset button
							if globals.Count == globals.ResetVal {
								gtx = gtx.Disabled()
							}

							for range c.resetBtn.Clicks() {
								globals.Count = globals.ResetVal
							}
							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["blue"],
									LabelColor: globals.Colours["white"],
									Button:     &c.resetBtn,
									Icon:       globals.RefreshIcon,
									Label:      strconv.FormatInt(globals.ResetVal, 10),
								}.Layout)
						}),

						globals.SpacerX,

						// Plus Button
						layout.Rigid(func(gtx C) D {
							for range c.plusBtn.Clicks() {
								globals.Count += globals.CountUnit
							}

							return globals.Inset.Layout(
								gtx,
								custom_widgets.LabeledIconBtn{
									Theme:      th,
									BgColor:    globals.Colours["green"],
									LabelColor: globals.Colours["black"],
									Button:     &c.plusBtn,
									Icon:       globals.PlusIcon,
									Label:      strconv.FormatInt(globals.CountUnit, 10),
								}.Layout,
							)
						}),
					)
				}),
			)
		}),
	)
}
