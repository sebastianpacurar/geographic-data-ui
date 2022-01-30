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
		leftList         layout.List
		checkboxes       []checkBox
		selectAllBoxes   widget.Clickable
		deselectAllBoxes widget.Clickable

		rightList layout.List
		radioBtns widget.Enum

		loaded bool
	}

	checkBox struct {
		name string
		widget.Bool
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
				Bool: widget.Bool{Value: table.ColState[table.ColNames[i]]},
			}
		}
		cds.loaded = true
	}

	txtSize := th.TextSize.Scale(12.0 / 15.0)
	btnInset := layout.Inset{
		Top: unit.Dp(5), Bottom: unit.Dp(5),
		Left: unit.Dp(6), Right: unit.Dp(6),
	}

	return layout.Flex{WeightSum: 2}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {

			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Bottom: unit.Dp(10)}.Layout(gtx,
						material.Body2(th, fmt.Sprintf("Columns (%d/%d)", getCheckedColumns(), len(table.ColState)-1)).Layout)
				}),

				layout.Rigid(func(gtx C) D {
					return layout.Inset{Top: unit.Dp(5), Bottom: unit.Dp(7.5)}.Layout(gtx, func(gtx C) D {
						switch {

						// check all unchecked checkboxes
						case cds.selectAllBoxes.Clicked():
							for i := range cds.checkboxes {
								if !cds.checkboxes[i].Value {
									cds.checkboxes[i].Value = true
									table.ColState[cds.checkboxes[i].name] = cds.checkboxes[i].Value
								}
							}
							op.InvalidateOp{}.Add(gtx.Ops)

						// uncheck all checkboxes, except "Capitals" checkbox
						case cds.deselectAllBoxes.Clicked():
							for i := range cds.checkboxes {
								if cds.checkboxes[i].name == table.CAPITALS {
									cds.checkboxes[i].Value = true
								} else {
									cds.checkboxes[i].Value = false
								}
								table.ColState[cds.checkboxes[i].name] = cds.checkboxes[i].Value

								// ignore if radioBtn Value is CAPITALS, since CAPITALS is the last default checked(true) checkbox
								if table.SearchBy != table.CAPITALS {
									// in case the column is unchecked, and the relative radioBtn Value is true, default it to NAME column
									table.SearchBy = table.NAME
									cds.radioBtns.Value = table.NAME
								}
							}
							op.InvalidateOp{}.Add(gtx.Ops)
						}
						return layout.Flex{}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								var btn material.ButtonStyle
								btn = material.Button(th, &cds.selectAllBoxes, "Select All")
								btn.TextSize = txtSize
								btn.Inset = btnInset
								btn.Background = globals.Colours[colours.LIGHT_SEA_GREEN]

								if len(table.ColState)-1 == getCheckedColumns() {
									gtx.Queue = nil
								}
								return btn.Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx C) D {
								var btn material.ButtonStyle
								btn = material.Button(th, &cds.deselectAllBoxes, "Deselect All")
								btn.TextSize = txtSize
								btn.Inset = btnInset
								btn.Background = globals.Colours[colours.FLAME_RED]

								// disable if there is one checked box in the list (prevents having all columns hidden)
								if getCheckedColumns() == 1 {
									gtx.Queue = nil
								}
								return btn.Layout(gtx)
							}),
						)
					})
				}),

				layout.Rigid(func(gtx C) D {
					return cds.leftList.Layout(gtx, len(table.ColNames), func(gtx C, i int) D {
						var cb material.CheckBoxStyle
						cb = material.CheckBox(th, &cds.checkboxes[i].Bool, cds.checkboxes[i].name)
						cb.Size = unit.Dp(18)
						cb.TextSize = txtSize
						cb.IconColor = globals.Colours[colours.LIGHT_SEA_GREEN]

						// invalidate in case it's the last marked box, so there is at least one column displayed
						if getCheckedColumns() == 1 && cds.checkboxes[i].Value {
							gtx.Queue = nil
						}

						if cb.CheckBox.Changed() {
							table.ColState[cds.checkboxes[i].name] = cb.CheckBox.Value
							//reorderColumns()
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

						// if the current column is unchecked then disable its radioBtn UI
						if !table.ColState[table.SearchByCols[i]] {
							// in case the column is unchecked, and the relative radioBtn Value is true, default it to NAME column
							if cds.radioBtns.Value == table.SearchByCols[i] {
								table.SearchBy = table.NAME
								cds.radioBtns.Value = table.NAME
								op.InvalidateOp{}.Add(gtx.Ops)
							}
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

//func reorderColumns() {
//	if table.SearchBy != table.NAME {
//		currVal := table.ColPos[table.SearchBy]
//
//		table.ColPos[table.SearchBy] = 0
//		for k, v := range table.ColPos {
//			if k != table.SearchBy {
//				if v < currVal {
//					table.ColPos[k] += 1
//				} else {
//					break
//				}
//			}
//		}
//	}
//}

// getCheckedColumns - returns the count of the displayed checked columns
func getCheckedColumns() int {
	count := 0
	for _, v := range table.ColState {
		if v {
			count++
		}
	}
	// return -1 since NAME column is sticky
	return count - 1
}
