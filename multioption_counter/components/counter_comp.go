package components

import (
	"gioui-experiment/multioption_counter/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"fmt"
	"strconv"
)

// Counter - component which controls the displayed data. It displays itself as:
// increase - this increments the current number.
// decrease - this decrements the current number.
// reset - this resets the current value to the ResetVal from globals.
type Counter struct {
	increase, decrease, reset widget.Clickable
}

func (c *Counter) Layout(th *material.Theme, gtx globals.C) globals.D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		layout.Flexed(1, func(gtx globals.C) globals.D {
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

		layout.Rigid(func(gtx globals.C) globals.D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(
				gtx,
				layout.Rigid(func(gtx globals.C) globals.D {
					return layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceEvenly,
					}.Layout(
						gtx,
						layout.Flexed(0.1, func(gtx globals.C) globals.D {
							for range c.increase.Clicks() {
								globals.Count++
							}
							btn := material.Button(th, &c.increase, "Increase")
							btn.Background = globals.Colours["green"]
							btn.Color = globals.Colours["black"]
							return btn.Layout(gtx)
						}),
						globals.SpacerX,
						layout.Flexed(0.1, func(gtx globals.C) globals.D {
							for range c.decrease.Clicks() {
								globals.Count--
							}
							btn := material.Button(th, &c.decrease, "Decrease")
							btn.Background = globals.Colours["red"]
							return btn.Layout(gtx)
						}),
					)
				}),

				globals.SpacerY,

				layout.Rigid(func(gtx globals.C) globals.D {
					for range c.reset.Clicks() {
						globals.Count = globals.ResetVal
					}
					btn := material.Button(th, &c.reset, fmt.Sprintf("Reset to %d", globals.ResetVal))
					btn.Background = globals.Colours["blue"]
					return btn.Layout(gtx)
				}),
			)
		}),
	)
}