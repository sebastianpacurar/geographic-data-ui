package components

import (
	"gioui-experiment/multioption_counter/globals"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"strconv"
	"strings"
)

type valueInput struct {
	widget.Editor
	invalid bool
	oldVal  string
}

type StartValue struct {
	valueInput
	changeVal widget.Clickable
}

func (vi *valueInput) updated() bool {
	newVal := vi.Editor.Text()
	updated := newVal != vi.oldVal
	vi.oldVal = newVal
	return updated
}

func (sv *StartValue) Layout(th *material.Theme, gtx C) D {
	editor := material.Editor(th, &sv.Editor, "Input incr/decr value")

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,
		layout.Flexed(1, func(gtx C) D {
			sv.Editor.SingleLine = true
			sv.Editor.Alignment = text.Middle
			editor.Font.Weight = text.Bold
			editor.TextSize = unit.Sp(20)
			border := widget.Border{
				Color:        globals.Colours["grey"],
				CornerRadius: unit.Dp(5),
				Width:        unit.Px(3),
			}
			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Sp(8)).Layout(
					gtx,
					editor.Layout,
				)
			})
		}),

		layout.Rigid(
			layout.Spacer{
				Width: globals.DefaultMargin,
			}.Layout,
		),

		layout.Flexed(1, func(gtx C) D {
			btn := material.Button(th, &sv.changeVal, "Change val to incr/decr")
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
	)
}
