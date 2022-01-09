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

		Click       widget.Clickable
		colList     layout.List
		loaded      bool
		ColumnNames []string
		Columns     []Cell
	}

	Cell struct {
		HeadCell  string
		Size      int
		IsEnabled bool
		Layout    func(C, *material.Theme, *Cell, color.NRGBA, bool) D
	}
)

func (r *Row) LayRow(gtx C, th *material.Theme, isHeader bool) D {
	rowColor := g.Colours[colours.ANTIQUE_WHITE]
	r.colList.Alignment = layout.Middle
	if !r.loaded {
		r.generateColumns()
		r.loaded = true
	}
	r.ColumnNames = []string{
		"Name", "Official Name", "Capital", "Region", "Subregion", "International Direct Dial Root", "Independent",
		"Status", "United Nations Member", "CCA 2", "CCA 3", "CCN 3", "Area", "Population",
	}

	border := widget.Border{
		Color: g.Colours[colours.GREY],
		Width: unit.Dp(1),
	}
	if isHeader {
		border.Width = unit.Dp(1.5)
		border.Color = g.Colours[colours.GREY]
	}

	return border.Layout(gtx, func(gtx C) D {
		if !isHeader {
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
				return r.colList.Layout(gtx, len(r.ColumnNames), func(gtx C, i int) D {
					return r.Columns[i].Layout(gtx, th, &r.Columns[i], rowColor, isHeader)
				})
			})
		} else {
			rowColor = g.Colours[colours.LAVENDERBLUSH]
			return r.colList.Layout(gtx, len(r.ColumnNames), func(gtx C, i int) D {
				return r.Columns[i].Layout(gtx, th, &r.Columns[i], rowColor, isHeader)
			})
		}
	})
}

func (r *Row) generateColumns() {
	for range r.ColumnNames {
		r.Columns = append(r.Columns,
			Cell{
				HeadCell: "Name",
				Size:     450,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.Name
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Official Name",
				Size:     550,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.OfficialName
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Capital",
				Size:     200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							capital := "N/A"
							if len(r.Capital) > 0 {
								capital = r.Capital[0]
							}
							if isHeader {
								capital = c.HeadCell
							}
							return material.Body1(th, capital).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Region",
				Size:     175,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.Region
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Subregion",
				Size:     200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.Subregion
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "International Direct Dial Root",
				Size:     280,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.IddRoot
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Independent",
				Size:     180,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							independent := "No"
							if r.Independent {
								independent = "Yes"
							}
							if isHeader {
								independent = c.HeadCell
							}
							return material.Body1(th, independent).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Status",
				Size:     175,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.Status
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "United Nations Member",
				Size:     200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							unMember := "No"
							if r.UNMember {
								unMember = "Yes"
							}
							if isHeader {
								unMember = c.HeadCell
							}
							return material.Body1(th, unMember).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "CCA 2",
				Size:     85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.Cca2
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "CCA 3",
				Size:     85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.Cca3
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "CCN 3",
				Size:     85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.Ccn3
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Area",
				Size:     100,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := fmt.Sprintf("%.3f", r.Area)
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Population",
				Size:     100,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := fmt.Sprintf("%d", int(r.Population))
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			})
	}
}
