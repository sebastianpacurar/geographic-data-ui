package components

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type TextEditor struct {
	in        textInput
	formatBtn widget.Clickable
}

type textInput struct {
	component.TextField
	isDisabled bool
	border     widget.Border
}

// InitTextFields - this sets the initial state of the fields
func (te *TextEditor) InitTextFields() {
	te.in.Editor.SingleLine = false
	te.in.Alignment = layout.Alignment(text.Start)
	te.in.isDisabled = false
}

// Layout - Still in progress!
func (te *TextEditor) Layout(th *material.Theme, gtx C) D {
	input := material.Editor(th, &te.in.Editor, "Type your Thoughts...")

	return layout.UniformInset(globals.DefaultMargin).Layout(gtx, func(gtx C) D {
		return layout.Flex{
			Axis:      layout.Horizontal,
			WeightSum: 1,
		}.Layout(
			gtx,
			// Input Text
			layout.Flexed(1, func(gtx C) D {
				input.TextSize = unit.Sp(20)
				input.HintColor = globals.Colours["dark-slate-grey"]
				border := widget.Border{
					Color:        globals.Colours["grey"],
					CornerRadius: unit.Dp(5),
					Width:        unit.Px(2),
				}
				switch {
				case te.in.Editor.Focused():
					border.Color = th.Palette.ContrastBg
					border.Width = unit.Px(3)
				}
				return border.Layout(gtx, func(gtx C) D {
					return layout.UniformInset(unit.Dp(8)).Layout(
						gtx,
						input.Layout,
					)
				})
			}),
		)
	})
}
