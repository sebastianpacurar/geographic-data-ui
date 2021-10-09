package components

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"fmt"
	"strconv"
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

		// Display the Counter number
		layout.Flexed(1, func(gtx C) D {
			currVal := material.H2(th, strconv.FormatInt(globals.Count, 10))
			if globals.Count < 0 {
				currVal.Color = globals.Colours["red"]
			} else if globals.Count > 0 {
				currVal.Color = globals.Colours["dark-green"]
			} else {
				currVal.Color = globals.Colours["grey"]
			}
			return layout.Center.Layout(
				gtx,
				currVal.Layout,
			)
		}),

		globals.SpacerY,

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
						layout.Flexed(0.1, func(gtx C) D {
							for range c.minusBtn.Clicks() {
								globals.Count--
							}
							btn := component.SimpleIconButton(
								globals.Colours["red"],
								globals.Colours["white"],
								&c.minusBtn,
								globals.MinusIcon,
							)
							return btn.Layout(gtx)
						}),

						globals.SpacerX,

						// Plus Button
						layout.Flexed(0.1, func(gtx C) D {
							for range c.plusBtn.Clicks() {
								globals.Count++
							}
							btn := component.SimpleIconButton(
								globals.Colours["green"],
								globals.Colours["black"],
								&c.plusBtn,
								globals.PlusIcon,
							)
							return btn.Layout(gtx)
						}),
					)
				}),

				globals.SpacerY,

				// Reset Button
				layout.Rigid(func(gtx C) D {
					// if count == reset, disable Reset button
					if globals.Count == globals.ResetVal {
						gtx = gtx.Disabled()
					}

					for range c.resetBtn.Clicks() {
						globals.Count = globals.ResetVal
					}
					btn := material.Button(th, &c.resetBtn, fmt.Sprintf("Reset to %d", globals.ResetVal))
					btn.Background = globals.Colours["blue"]
					return btn.Layout(gtx)
				}),
			)
		}),
	)
}
