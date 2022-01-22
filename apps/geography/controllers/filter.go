package controllers

import (
	"fmt"
	"gioui-experiment/apps/geography/table"
	"gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	FilterTable struct {
		leftList   layout.List
		checkboxes []checkBox

		rightList layout.List
		radioBtns widget.Enum

		loaded bool
	}

	checkBox struct {
		name string
		box  widget.Bool
	}
)

// Layout - Lays out the column checkboxes
func (ft *FilterTable) Layout(gtx C, th *material.Theme) D {
	if !ft.loaded {
		ft.leftList.Axis = layout.Vertical
		ft.rightList.Axis = layout.Vertical
		ft.radioBtns.Value = table.SearchBy
		ft.checkboxes = make([]checkBox, len(table.ColNames))
		for i := range ft.checkboxes {
			ft.checkboxes[i] = checkBox{
				name: table.ColNames[i],
				box:  widget.Bool{Value: table.ColState[table.ColNames[i]]},
			}
		}
		ft.loaded = true
	}

	return layout.Flex{WeightSum: 2}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {

			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Bottom: unit.Dp(10)}.Layout(gtx,
						material.Body2(th, fmt.Sprintf("Columns (%d/%d)", ft.getCheckedColumns(), len(table.ColState))).Layout)
				}),

				layout.Flexed(1, func(gtx C) D {
					return ft.leftList.Layout(gtx, len(table.ColNames), func(gtx C, i int) D {
						var cb material.CheckBoxStyle
						cb = material.CheckBox(th, &ft.checkboxes[i].box, ft.checkboxes[i].name)
						cb.Size = unit.Dp(18)
						cb.TextSize = th.TextSize.Scale(12.0 / 15.0)
						cb.IconColor = globals.Colours[colours.LIGHT_SEA_GREEN]

						if cb.CheckBox.Changed() {
							table.ColState[ft.checkboxes[i].name] = cb.CheckBox.Value
							op.InvalidateOp{}.Add(gtx.Ops)
						}
						return cb.Layout(gtx)
					})
				}),
			)

		}),
		layout.Flexed(1, func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Bottom: unit.Dp(10)}.Layout(gtx,
						material.Body2(th, "Search by").Layout)
				}),

				layout.Flexed(1, func(gtx C) D {
					if ft.radioBtns.Changed() {
						table.SearchBy = ft.radioBtns.Value
						op.InvalidateOp{}.Add(gtx.Ops)
					}

					return ft.rightList.Layout(gtx, len(table.SearchByCols), func(gtx C, i int) D {
						var rbtn material.RadioButtonStyle
						rbtn = material.RadioButton(th, &ft.radioBtns, table.SearchByCols[i], table.SearchByCols[i])

						rbtn.Size = unit.Dp(18)
						rbtn.TextSize = th.TextSize.Scale(12.0 / 15.0)
						rbtn.IconColor = globals.Colours[colours.LIGHT_SEA_GREEN]

						if !table.ColState[table.SearchByCols[i]] {
							rbtn.Color = globals.Colours[colours.FLAME_RED]
							gtx.Queue = nil
						}

						return rbtn.Layout(gtx)
					})
				}),
			)
		}),
	)
}

func (ft *FilterTable) getCheckedColumns() int {
	count := 0
	for _, v := range table.ColState {
		if v {
			count++
		}
	}
	// return -1 since NAME column is sticky
	return count - 1
}
