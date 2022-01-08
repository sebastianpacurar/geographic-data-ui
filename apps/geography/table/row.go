package table

import (
	"fmt"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

type (
	Row struct {
		Click        widget.Clickable
		List         widget.List
		Name         string
		OfficialName string
		Capital      []string
		Independent  bool
		Status       string
		UNMember     bool
		Cca2         string
		Cca3         string
		Ccn3         string
		Area         float64
		Population   int32
		Region       string
		Subregion    string
		IddRoot      string
		Active       bool
		Selected     bool
		IsCPViewed   bool

		loaded      bool
		ColumnNames []string
		Columns     []Cell
	}

	Cell struct {
		HeadCell  string
		Size      int
		IsEnabled bool
		Layout    func(C, *material.Theme, *Cell, color.NRGBA) D
	}
)

func (r *Row) LayRow(gtx C, th *material.Theme) D {
	r.ColumnNames = []string{
		"Name",
		"Official Name",
		"Capital",
		"Region",
		"Subregion",
		"International Direct Dial Root",
		"Independent",
		"Status",
		"United Nations Member",
		"CCA 2",
		"CCA 3",
		"CCN 3",
		"Area",
		"Population",
	}

	var dims D

	border := widget.Border{
		Color: g.Colours[colours.GREY],
		Width: unit.Dp(1),
	}

	r.List.Axis = layout.Horizontal
	r.List.Alignment = layout.Middle
	if !r.loaded {
		r.generateColumns()
		r.loaded = true
	}

	rowColor := g.Colours[colours.ANTIQUE_WHITE]

	dims = material.List(th, &r.List).Layout(gtx, len(r.ColumnNames), func(gtx C, i int) D {
		return r.Columns[i].Layout(gtx, th, &r.Columns[i], rowColor)
	})

	return border.Layout(gtx, func(gtx C) D {
		return material.Clickable(gtx, &r.Click, func(gtx C) D {
			if r.Selected {
				rowColor = g.Colours[colours.AERO_BLUE]
			}
			if r.Click.Hovered() {
				if r.Selected {
					rowColor = g.Colours[colours.LIGHT_SALMON]
				} else {
					rowColor = g.Colours[colours.NYANZA]
				}
			}
			return dims
		})
	})
}

func (r *Row) generateColumns() {
	for i := range r.ColumnNames {
		r.Columns = append(r.Columns,
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     450,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.Name).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     550,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.OfficialName).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							capital := "N/A"
							if len(r.Capital) > 0 {
								capital = r.Capital[0]
							}
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, capital).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     175,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.Region).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.Subregion).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     280,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.IddRoot).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     180,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							independent := "No"
							if r.Independent {
								independent = "Yes"
							}
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, independent).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     175,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.Status).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							unMember := "No"
							if r.UNMember {
								unMember = "Yes"
							}
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, unMember).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.Cca2).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.Cca3).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, r.Ccn3).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     100,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),

						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, fmt.Sprintf("%.3f", r.Area)).Layout)
						}))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     100,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),

						layout.Stacked(func(gtx C) D {
							return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, material.Body1(th, fmt.Sprintf("%d", int(r.Population))).Layout)
						}))
				},
			})
	}
}
