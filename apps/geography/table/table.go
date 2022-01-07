package table

import (
	"gioui-experiment/apps/geography/data"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colours"
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
		rows       []Row
		rowList    widget.List
		columnList widget.List
		headerList widget.List
		loaded     bool
	}
)

func (t *Table) Layout(gtx C, th *material.Theme) D {
	t.rowList.Axis = layout.Vertical
	t.rowList.Alignment = layout.Middle
	t.columnList.Axis = layout.Horizontal
	t.columnList.Alignment = layout.Middle

	border := widget.Border{
		Color: g.Colours[colours.GREY],
		Width: unit.Dp(1),
	}

	if !t.loaded {
		for i := range data.Cached {
			t.rows = append(t.rows, Row{
				Name:         data.Cached[i].Name.Common,
				OfficialName: data.Cached[i].Name.Official,
				Capital:      data.Cached[i].Capital,
				Independent:  data.Cached[i].Independent,
				Status:       data.Cached[i].Status,
				UNMember:     data.Cached[i].UNMember,
				Cca2:         data.Cached[i].Cca2,
				Cca3:         data.Cached[i].Cca3,
				Ccn3:         data.Cached[i].Ccn3,
				Area:         data.Cached[i].Area,
				Population:   data.Cached[i].Population,
				Region:       data.Cached[i].Region,
				Subregion:    data.Cached[i].Subregion,
				IddRoot:      data.Cached[i].Idd.Root,
				Active:       data.Cached[i].Active,
			})
		}
		t.loaded = true
	} else {
		for i := range data.Cached {
			t.rows[i].Active = data.Cached[i].Active
			t.rows[i].Selected = data.Cached[i].Selected
			t.rows[i].IsCPViewed = data.Cached[i].IsCPViewed
		}
	}

	return material.List(th, &t.columnList).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,

			// Header Area
			layout.Rigid(func(gtx C) D {
				return layout.Flex{}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return border.Layout(gtx, func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(2),
								Bottom: unit.Dp(2),
								Left:   unit.Dp(2),
							}.Layout(gtx, func(gtx C) D {
								return layout.Stack{Alignment: layout.Center}.Layout(gtx,

									layout.Expanded(func(gtx C) D {
										return g.ColoredArea(gtx, image.Pt(450, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "Name").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(550, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "Official Name").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(200, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "Capital").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(175, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "Region").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(220, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "Subregion").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(280, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "International Direct Dial Root").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(180, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "independent").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(175, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "Status").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(200, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, "United Nations Member").Layout(gtx)
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
										return g.ColoredArea(gtx, image.Pt(85, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(material.Body1(th, "CCA 2").Layout))
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
										return g.ColoredArea(gtx, image.Pt(85, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(material.Body1(th, "CCA 3").Layout))
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
										return g.ColoredArea(gtx, image.Pt(85, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(material.Body1(th, "CCN 3").Layout))
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
										return g.ColoredArea(gtx, image.Pt(100, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(material.Body1(th, "Area").Layout))
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
										return g.ColoredArea(gtx, image.Pt(100, gtx.Constraints.Min.Y), g.Colours[colours.ANTIQUE_WHITE])
									}),

									layout.Stacked(material.Body1(th, "Population").Layout))
							})
						})
					}))
			}),

			// Row Area
			layout.Flexed(1, func(gtx C) D {
				return material.List(th, &t.rowList).Layout(gtx, len(data.Cached), func(gtx C, i int) D {
					var dims D
					if t.rows[i].Active {
						if t.rows[i].Click.Clicked() {
							if t.rows[i].Selected {
								data.Cached[i].Selected = false
							} else {
								data.Cached[i].Selected = true
							}
						}
						dims = t.rows[i].LayRow(gtx, th)
					}
					return dims
				})
			}))
	})
}
