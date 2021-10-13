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

type ValueHandler struct {
	startVal, unitVal       component.TextField
	changeStart, changeUnit widget.Clickable
}

// InitTextFields - initializes the state for the TextFields
func (vh *ValueHandler) InitTextFields() {
	vh.startVal.SingleLine = true
	vh.unitVal.SingleLine = true
	vh.startVal.Alignment = layout.Alignment(text.Start)
	vh.unitVal.Alignment = layout.Alignment(text.Start)
}

// Layout - displays the Start From button and text input horizontally.
// The editor is flexed, so it can enlarge/shrink while resizing on the X-Axis.
func (vh *ValueHandler) Layout(th *material.Theme, gtx C) D {
	eStart := material.Editor(th, &vh.startVal.Editor, "0")
	eUnit := material.Editor(th, &vh.unitVal.Editor, "0")

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,

		// "Start From" button - to enable the changed start value

		/////
		layout.Rigid(func(gtx C) D {
			btn := material.Button(th, &vh.changeStart, "Start From")
			btn.Background = globals.Colours["blue"]
			btn.Color = globals.Colours["white"]

			switch {
			case isFieldNumeric(vh.startVal):
				gtx = gtx.Disabled()
			case vh.changeStart.Clicked():
				inpVal := vh.startVal.Text()
				inpVal = strings.TrimSpace(inpVal)
				intVal, _ := strconv.ParseInt(inpVal, 10, 64)
				vh.startVal.Clear()
				globals.Count = intVal
				globals.ResetVal = intVal
			}
			return btn.Layout(gtx)
		}),

		globals.SpacerX,

		// TextField Widget - to change start value
		layout.Rigid(func(gtx C) D {
			eStart.TextSize = unit.Sp(20)
			eStart.HintColor = globals.Colours["dark-slate-grey"]
			border := globals.DefaultBorder
			validateTextField(th, vh.startVal, eStart, &border)

			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(
					gtx,
					eStart.Layout,
				)
			})
		}),

		///////

		layout.Rigid(func(gtx C) D {
			if vh.changeUnit.Clicked() {
				inpVal := vh.unitVal.Text()
				inpVal = strings.TrimSpace(inpVal)
				intVal, _ := strconv.ParseInt(inpVal, 10, 64)
				globals.CountUnit = intVal
			}
			btn := material.Button(th, &vh.changeUnit, "Set Unit To")
			btn.Background = globals.Colours["blue"]
			return btn.Layout(gtx)
		}),

		globals.SpacerX,

		layout.Rigid(func(gtx C) D {
			eUnit.TextSize = unit.Sp(20)
			eUnit.HintColor = globals.Colours["dark-slate-grey"]
			border := globals.DefaultBorder
			validateTextField(th, vh.startVal, eStart, &border)

			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(
					gtx,
					eUnit.Layout,
				)
			})
		}),
	)
}

// isFieldNumeric - verifies if the string can be converted to integer.
// It iterates through the input and verifies if each rune digit casted to string,
// can be converted to integer. Exception from the rule is the minus sign as first char.
func isFieldNumeric(e component.TextField) bool {
	if e.Len() > 0 {
		for i := range e.Text() {
			if i == 0 && e.Text()[0] == '-' {
				continue
			}
			_, err := strconv.Atoi(string(e.Text()[i]))
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
func trimInput(e component.TextField, count int) {
	str := e.Text()
	for i := 0; i < count; i++ {
		_, _ = utf8.DecodeLastRuneInString(str)
	}
	e.SetText(str[:len(str)-count])
}

// 1) In case the length of the input is >= 18, then keep trimming chars
// until the length is back to 17
// 2) If there are non-numeric characters, change colors to red
// 3) When Focused change border color a bit
func validateTextField(th *material.Theme, e component.TextField, eStyle material.EditorStyle, b *widget.Border) {
	switch {
	case e.Len() >= 18:
		trimInput(e, e.Len()-18)
	case isFieldNumeric(e):
		b.Color = globals.Colours["red"]
		b.Width = unit.Px(5)
		eStyle.Color = globals.Colours["dark-red"]
	case e.Focused():
		b.Color = th.Palette.ContrastBg
		b.Width = unit.Px(3)
	}
}
