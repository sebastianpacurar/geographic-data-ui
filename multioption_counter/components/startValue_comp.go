package components

import (
	"gioui-experiment/multioption_counter/globals"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"strconv"
	"strings"
)

// valueInput - this is the Editor of the app
// TODO: switch to TextField from gioui.or/x/
type valueInput struct {
	component.TextField
}

// StartValue - component that keeps track of its state.
// It displays itself as an Editor and a Button
type StartValue struct {
	valueInput
	changeVal widget.Clickable
}

// Layout - displays the Start From button and text input horizontally.
// The editor is flexed, so it can enlarge/shrink while resizing on the X-Axis.
func (sv *StartValue) Layout(th *material.Theme, gtx globals.C) globals.D {
	editor := material.Editor(th, &sv.Editor, "0")

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,

		// "Start From" button - to enable the changed start value
		layout.Rigid(func(gtx globals.C) globals.D {
			btn := material.Button(th, &sv.changeVal, "Start From")
			btn.Background = globals.Colours["blue"]
			if sv.changeVal.Clicked() {
				inpVal := sv.Text()
				inpVal = strings.TrimSpace(inpVal)
				intVal, _ := strconv.ParseInt(inpVal, 10, 64)
				globals.Count = intVal
				globals.ResetVal = intVal
			}
			return btn.Layout(gtx)
		}),

		globals.SpacerX,

		// Input widget - to change start value
		layout.Flexed(1, func(gtx globals.C) globals.D {
			sv.Editor.SingleLine = true
			sv.Editor.Alignment = text.Middle
			editor.Font.Weight = text.Bold
			editor.TextSize = unit.Sp(20)
			border := widget.Border{
				Color:        globals.Colours["grey"],
				CornerRadius: unit.Dp(5),
				Width:        unit.Px(3),
			}
			return border.Layout(gtx, func(gtx globals.C) globals.D {
				return layout.UniformInset(unit.Sp(8)).Layout(
					gtx,
					editor.Layout,
				)
			})
		}),
	)
}
