package components

import (
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type TextArea struct {
	field  widget.Editor
	border widget.Border
}

func (ta *TextArea) Layout(th *material.Theme) layout.FlexChild {
	input := material.Editor(th, &ta.field, "Type your Thoughts...")
	input.SelectionColor = g.Colours[colors.TEXT_SELECTION]
	ta.field.SingleLine = false
	ta.field.Alignment = text.Start

	return layout.Flexed(1, func(gtx C) D {
		border := widget.Border{
			Color:        g.Colours[colors.GREY],
			CornerRadius: unit.Dp(5),
			Width:        unit.Px(2),
		}
		switch {
		case ta.field.Focused():
			border.Color = th.Palette.ContrastBg
			border.Width = unit.Px(2)
		}

		return g.Inset.Layout(gtx, func(gtx C) D {
			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
					return input.Layout(gtx)
				})
			})
		})
	})
}
