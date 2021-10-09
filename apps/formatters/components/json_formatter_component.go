package components

import (
	"encoding/json"
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

type JsonFormatter struct {
	in        textInput
	out       textOutput
	formatBtn widget.Clickable
}

type textInput struct {
	component.TextField
	isDisabled bool
	border     widget.Border
}

type textOutput struct {
	component.TextField
	isDisabled bool
	border     widget.Border
}

// InitTextFields - this sets the initial state of the fields
func (jf *JsonFormatter) InitTextFields() {
	jf.in.Editor.SingleLine = false
	jf.in.Alignment = layout.Alignment(text.Start)
	jf.in.isDisabled = false

	jf.out.Editor.SingleLine = false
	jf.out.Alignment = layout.Alignment(text.Start)
	jf.out.isDisabled = true
}

func (jf *JsonFormatter) Layout(th *material.Theme, gtx C) D {
	input := material.Editor(th, &jf.in.Editor, "Paste json here...")
	output := material.Editor(th, &jf.out.Editor, "Click GO! button and see magic...")
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		// Input Text
		layout.Flexed(1, func(gtx C) D {
			input.TextSize = unit.Sp(20)
			input.HintColor = globals.Colours["dark-slate-grey"]
			border := widget.Border{
				Color:        globals.Colours["grey"],
				CornerRadius: unit.Dp(5),
				Width:        unit.Px(3),
			}
			switch {
			case jf.in.Editor.Focused():
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

		globals.SpacerY,

		// Output text
		layout.Flexed(1, func(gtx C) D {
			output.TextSize = unit.Sp(20)
			output.HintColor = globals.Colours["grey"]
			if jf.out.Len() == 0 {
				gtx = gtx.Disabled()
			}
			border := widget.Border{
				Color:        globals.Colours["black"],
				CornerRadius: unit.Dp(5),
				Width:        unit.Px(2),
			}
			switch {
			case jf.out.Editor.Focused():
				border.Color = th.Palette.ContrastBg
				border.Width = unit.Px(3)
			case jf.out.isDisabled:
				border.Color = globals.Colours["red"]
				border.Width = unit.Px(5)
				output.Hint = "Please add a valid json string!"
				output.HintColor = globals.Colours["dark-red"]
			}
			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(
					gtx,
					output.Layout,
				)
			})
		}),
	)
}

// isInJsonString - checks to see if the provided Input is a JSON string
func (jf *JsonFormatter) isInJsonString(str string) bool {
	var jsonMessage json.RawMessage
	err := json.Unmarshal([]byte(str), &jsonMessage)
	if err != nil {
		return false
	}
	return true
}

// isInJson - checks to see if the provided input is a JSON (interface)
func (jf *JsonFormatter) isInJson(str string) bool {
	var jsonMessage map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonMessage)
	if err != nil {
		return false
	}
	return true
}
