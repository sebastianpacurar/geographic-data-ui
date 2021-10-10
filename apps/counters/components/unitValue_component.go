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
)

type UnitVal struct {
	component.TextField
	enableBtn widget.Clickable
}

func (u *UnitVal) InitTextField() {
	u.SingleLine = true
	u.Alignment = layout.Alignment(text.Start)
}

func (u *UnitVal) Layout(th *material.Theme, gtx C) D {
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx C) D {
			if u.enableBtn.Clicked() {
				inpVal := u.Text()
				inpVal = strings.TrimSpace(inpVal)
				intVal, _ := strconv.ParseInt(inpVal, 10, 64)
				globals.CountUnit = intVal
			}
			btn := material.Button(th, &u.enableBtn, "Set Unit To")
			btn.Background = globals.Colours["blue"]
			return btn.Layout(gtx)
		}),

		globals.SpacerX,

		layout.Rigid(func(gtx C) D {
			editor := material.Editor(th, &u.Editor, "1")
			editor.TextSize = unit.Sp(20)
			editor.HintColor = globals.Colours["dark-slate-grey"]
			border := widget.Border{
				Color:        globals.Colours["grey"],
				CornerRadius: unit.Dp(3),
				Width:        unit.Px(2),
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
