package components

import (
	"gioui-experiment/multioption_counter/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"strconv"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Counter struct {
	increase, decrease, reset widget.Clickable
}

func (c *Counter) Layout(th *material.Theme, gtx C) D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
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

		layout.Rigid(
			layout.Spacer{
				Height: globals.DefaultMargin,
			}.Layout,
		),

		layout.Flexed(0.1, func(gtx C) D {
			for range c.increase.Clicks() {
				globals.Count++
			}
			btn := material.Button(th, &c.increase, "Increase")
			btn.Background = globals.Colours["green"]
			btn.Color = globals.Colours["black"]
			return btn.Layout(gtx)
		}),

		layout.Rigid(
			layout.Spacer{
				Height: globals.DefaultMargin,
			}.Layout,
		),

		layout.Flexed(0.1, func(gtx C) D {
			for range c.decrease.Clicks() {
				globals.Count--
			}
			btn := material.Button(th, &c.decrease, "Decrease")
			btn.Background = globals.Colours["red"]
			return btn.Layout(gtx)
		}),

		layout.Rigid(
			layout.Spacer{
				Height: globals.DefaultMargin,
			}.Layout,
		),

		layout.Flexed(0.1, func(gtx C) D {
			for range c.reset.Clicks() {
				globals.Count = globals.ResetVal
			}
			btn := material.Button(th, &c.reset, "Reset")
			btn.Background = globals.Colours["blue"]
			return btn.Layout(gtx)
		}),
	)
}
