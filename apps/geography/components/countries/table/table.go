package table

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Table struct {
		rows       []row
		rowList    widget.List
		columnList widget.List
		loaded     bool
	}

	row struct {
		click        widget.Clickable
		name         string
		officialName string
		capital      []string
		independent  bool
		status       string
		unMember     bool
		cca2         string
		cca3         string
		ccn3         string
		area         float64
		population   int32
		active       bool
		selected     bool
		crossSize    int
	}
)

func (t *Table) Layout(gtx C, th *material.Theme) D {
	t.rowList.Axis = layout.Vertical
	t.rowList.Alignment = layout.Middle
	t.columnList.Axis = layout.Horizontal
	t.columnList.Alignment = layout.Middle

	if !t.loaded {
		for i := range data.Data {
			t.rows = append(t.rows, row{
				name:         data.Data[i].Name.Common,
				officialName: data.Data[i].Name.Official,
				capital:      data.Data[i].Capital,
				independent:  data.Data[i].Independent,
				status:       data.Data[i].Status,
				unMember:     data.Data[i].UNMember,
				cca2:         data.Data[i].Cca2,
				cca3:         data.Data[i].Cca3,
				ccn3:         data.Data[i].Ccn3,
				area:         data.Data[i].Area,
				population:   data.Data[i].Population,
				active:       data.Data[i].Active,
			})
		}
		t.loaded = true
	} else {
		for i := range data.Data {
			t.rows[i].active = data.Data[i].Active
		}
	}
	border := widget.Border{
		Color: g.Colours[colors.GREY],
		Width: unit.Px(1),
	}

	return material.List(th, &t.columnList).Layout(gtx, 1, func(gtx C, _ int) D {
		return material.List(th, &t.rowList).Layout(gtx, len(data.Data), func(gtx C, i int) D {
			return material.Clickable(gtx, &t.rows[i].click, func(gtx C) D {
				var content D
				rowColor := g.Colours[colors.ANTIQUE_WHITE]

				if t.rows[i].selected {
					rowColor = g.Colours[colors.AERO_BLUE]
				}

				if t.rows[i].active {
					if t.rows[i].click.Clicked() {
						if t.rows[i].selected {
							t.rows[i].selected = false
						} else {
							t.rows[i].selected = true
						}
					}

					if t.rows[i].click.Hovered() {
						data.Data[i].Hovered = true
						if !t.rows[i].selected {
							rowColor = g.Colours[colors.NYANZA]
						} else {
							rowColor = g.Colours[colors.LIGHT_SALMON]
						}
					} else {
						data.Data[i].Hovered = false
					}

					content = layout.Flex{}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(450, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(func(gtx C) D {
											return material.Body1(th, t.rows[i].name).Layout(gtx)
										}))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(550, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(func(gtx C) D {
											return material.Body1(th, t.rows[i].officialName).Layout(gtx)
										}))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(200, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(func(gtx C) D {
											capital := "N/A"
											if len(t.rows[i].capital) > 0 {
												capital = t.rows[i].capital[0]
											}
											return material.Body1(th, capital).Layout(gtx)
										}))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(func(gtx C) D {
											independent := "No"
											if t.rows[i].independent {
												independent = "Yes"
											}
											return material.Body1(th, independent).Layout(gtx)
										}))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(175, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(func(gtx C) D {
											return material.Body1(th, t.rows[i].status).Layout(gtx)
										}))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(func(gtx C) D {
											unMember := "No"
											if t.rows[i].unMember {
												unMember = "Yes"
											}
											return material.Body1(th, unMember).Layout(gtx)
										}))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(material.Body1(th, t.rows[i].cca2).Layout))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(material.Body1(th, t.rows[i].cca3).Layout))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(material.Body1(th, t.rows[i].ccn3).Layout))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(100, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(material.Body1(th, fmt.Sprintf("%.0f", t.rows[i].area)).Layout))
								})
							})
						}),
						layout.Rigid(func(gtx C) D {
							return border.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(2),
									Bottom: unit.Dp(2),
									Left:   unit.Dp(2),
								}.Layout(gtx, func(gtx C) D {
									return layout.Stack{Alignment: layout.Center}.Layout(gtx,

										layout.Expanded(func(gtx C) D {
											return g.ColoredArea(gtx, image.Pt(100, gtx.Constraints.Min.Y), rowColor)
										}),

										layout.Stacked(material.Body1(th, fmt.Sprintf("%d", int(t.rows[i].population))).Layout))
								})
							})
						}),
					)
				}
				return content
			})
		})
	})
}
