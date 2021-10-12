package components

import (
	"gioui-experiment/custom_widgets"
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"strconv"
)

var (
//plusBtn = new(widget.Clickable)
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Counter - component which controls the displayed data. It displays itself as:
// plusBtn - this increments the current number.
// minusBtn - this decrements the current number.
// resetBtn - this resets the current value to the ResetVal from globals.
type Counter struct {
	plusBtn, minusBtn, resetBtn widget.Clickable
}

// Layout - returns the Dimensions for the Counter design structure:
// plusBtn, minusBtn = 2 rounded SimpleIconButtons
// resetBtn is below them and fills the whole width
func (c *Counter) Layout(th *material.Theme, gtx C) D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,

		// Using flex-box on Y-Axis, which contains the 3 buttons
		// The first child of the Flex layout contains a Rigid layout, in which the 2 buttons are flexed horizontally.
		// The second child is the reset Button, returned as a layout.Rigid
		layout.Rigid(func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(
				gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceEvenly,
					}.Layout(
						gtx,

						// Minus Button
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
									BgColor:    globals.Colours["red"],
									LabelColor: globals.Colours["white"],
									Button:     &c.minusBtn,
									Icon:       globals.MinusIcon,
									Label:      strconv.FormatInt(globals.CountUnit, 10),
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
