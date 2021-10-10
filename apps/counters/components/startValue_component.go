package components

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"strconv"
	"strings"
	"unicode/utf8"
)

// StartValue - component that keeps track of its state.
// It displays itself as an Editor and a Button
type StartValue struct {
	component.TextField
	changeVal widget.Clickable
}

// InitTextField - initializes the state for the TextField
func (sv *StartValue) InitTextField() {
	sv.SingleLine = true
	sv.Alignment = layout.Alignment(text.Start)
}

// isFieldNumeric - verifies if the string can be converted to integer.
// It iterates through the input and verifies if each rune digit casted to string,
// can be converted to integer. Exception from the rule is the minus sign as first char.
func (sv *StartValue) isFieldNumeric() bool {
	if sv.Len() > 0 {
		for i := range sv.Text() {
			if i == 0 && sv.Text()[0] == '-' {
				continue
			}
			_, err := strconv.Atoi(string(sv.Text()[i]))
			if err != nil {
				return false
			}
		}
	}
	return true
}

// trimInput - removes the last rune from the string iteratively. Useful in case
// the input char count passes 17 chars, and mostly In case the system breaks and
// more than 18 characters can be provided.
func (sv *StartValue) trimInput(count int) {
	str := sv.Text()
	for i := 0; i < count; i++ {
		_, _ = utf8.DecodeLastRuneInString(str)
	}
	sv.SetText(str[:len(str)-count])
}

// Layout - displays the Start From button and text input horizontally.
// The editor is flexed, so it can enlarge/shrink while resizing on the X-Axis.
func (sv *StartValue) Layout(th *material.Theme, gtx C) D {
	editor := material.Editor(th, &sv.Editor, "0")

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,

		// "Start From" button - to enable the changed start value
		layout.Rigid(func(gtx C) D {
			btn := material.Button(th, &sv.changeVal, "Start From")
			btn.Background = globals.Colours["blue"]
			btn.Color = globals.Colours["white"]

			switch {
			case !sv.isFieldNumeric():
				gtx = gtx.Disabled()
			case sv.changeVal.Clicked():
				inpVal := sv.Text()
				inpVal = strings.TrimSpace(inpVal)
				intVal, _ := strconv.ParseInt(inpVal, 10, 64)
				sv.Clear()
				globals.Count = intVal
				globals.ResetVal = intVal
			}
			return btn.Layout(gtx)
		}),

		globals.SpacerX,

		// TextField Widget - to change start value
		layout.Rigid(func(gtx C) D {
			editor.TextSize = unit.Sp(20)
			editor.HintColor = globals.Colours["dark-slate-grey"]
			border := globals.DefaultBorder

			// 1) In case the length of the input is >= 18, then keep trimming chars
			// until the length is back to 17
			// 2) If there are non-numeric characters, change colors to red
			// 3) When Focused change border color a bit
			switch {
			case sv.Len() >= 18:
				sv.trimInput(sv.Len() - 18)
			case !sv.isFieldNumeric():
				border.Color = globals.Colours["red"]
				border.Width = unit.Px(5)
				editor.Color = globals.Colours["dark-red"]
			case sv.Focused():
				border.Color = th.Palette.ContrastBg
				border.Width = unit.Px(3)
			}
			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(
					gtx,
					editor.Layout,
				)
			})
		}),
	)
}
