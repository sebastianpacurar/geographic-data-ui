package components

import (
	color "gioui-experiment/custom_themes/colors"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type TextArea struct {
	field  widget.Editor
	border widget.Border
}

func (ta *TextArea) Layout(gtx C, th *material.Theme) D {
	input := material.Editor(th, &ta.field, "Type your Thoughts...")
	input.SelectionColor = g.Colours[color.TEXT_SELECTION]
	ta.field.SingleLine = false
	ta.field.Alignment = text.Start

	textArea := layout.Flexed(1, func(gtx C) D {
		border := widget.Border{
			Color:        g.Colours[color.GREY],
			CornerRadius: unit.Dp(5),
			Width:        unit.Px(2),
		}
		switch {
		case ta.field.Focused():
			border.Color = th.Palette.ContrastBg
			border.Width = unit.Px(3)
		}
		return border.Layout(gtx, func(gtx C) D {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
				return input.Layout(gtx)
			})
		})
	})
	return layout.Flex{}.Layout(gtx, textArea)
}
