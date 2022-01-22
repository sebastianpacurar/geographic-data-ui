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

	ColDisplaySearch struct {
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
func (cds *ColDisplaySearch) Layout(gtx C, th *material.Theme) D {
	if !cds.loaded {
		cds.leftList.Axis = layout.Vertical
		cds.rightList.Axis = layout.Vertical
		cds.radioBtns.Value = table.SearchBy
		cds.checkboxes = make([]checkBox, len(table.ColNames))
		for i := range cds.checkboxes {
			cds.checkboxes[i] = checkBox{
				name: table.ColNames[i],
				box:  widget.Bool{Value: table.ColState[table.ColNames[i]]},
			}
		}
		cds.loaded = true
	}

	return layout.Flex{WeightSum: 2}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {

			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Bottom: unit.Dp(10)}.Layout(gtx,
						material.Body2(th, fmt.Sprintf("Columns (%d/%d)", cds.getCheckedColumns(), len(table.ColState))).Layout)
				}),

				layout.Rigid(func(gtx C) D {
					return cds.leftList.Layout(gtx, len(table.ColNames), func(gtx C, i int) D {
						var cb material.CheckBoxStyle
						cb = material.CheckBox(th, &cds.checkboxes[i].box, cds.checkboxes[i].name)
						cb.Size = unit.Dp(18)
						cb.TextSize = th.TextSize.Scale(12.0 / 15.0)
						cb.IconColor = globals.Colours[colours.LIGHT_SEA_GREEN]

						if cb.CheckBox.Changed() {
							table.ColState[cds.checkboxes[i].name] = cb.CheckBox.Value
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

				layout.Rigid(func(gtx C) D {
					if cds.radioBtns.Changed() {
						table.SearchBy = cds.radioBtns.Value
						op.InvalidateOp{}.Add(gtx.Ops)
					}

					return cds.rightList.Layout(gtx, len(table.SearchByCols), func(gtx C, i int) D {
						var rbtn material.RadioButtonStyle
						rbtn = material.RadioButton(th, &cds.radioBtns, table.SearchByCols[i], table.SearchByCols[i])

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

// getCheckedColumns - returns the count of the displayed checked columns
func (cds *ColDisplaySearch) getCheckedColumns() int {
	count := 0
	for _, v := range table.ColState {
		if v {
			count++
		}
	}
	// return -1 since NAME column is sticky
	return count - 1
}
