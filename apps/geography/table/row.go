package table

import (
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
		//"Region",
		//"Subregion",
		//"International Direct Dial Root",
		//"Independent",
		//"Status",
		//"United Nations Member",
		//"CCA 2",
		//"CCA 3",
		//"CCN3",
	}

	var dims D

	r.List.Axis = layout.Horizontal
	r.List.Alignment = layout.Middle
	if !r.loaded {
		r.generateColumns()
		r.loaded = true
	}
	border := widget.Border{
		Color: g.Colours[colours.GREY],
		Width: unit.Dp(1),
	}
	rowColor := g.Colours[colours.ANTIQUE_WHITE]

	dims = material.List(th, &r.List).Layout(gtx, len(r.ColumnNames), func(gtx C, i int) D {
		return border.Layout(gtx, func(gtx C) D {
			return r.Columns[i].Layout(gtx, th, &r.Columns[i], rowColor)
		})
	})

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
							return g.ColoredArea(gtx, image.Pt(c.Size, 80), color)
						}),
						layout.Stacked(material.Body1(th, r.Name).Layout))
				},
			},
			Cell{
				HeadCell: r.ColumnNames[i],
				Size:     550,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, 80), color)
						}),
						layout.Stacked(material.Body1(th, r.OfficialName).Layout))
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
							Capital := "N/A"
							if len(r.Capital) > 0 {
								Capital = r.Capital[0]
							}
							return material.Body1(th, Capital).Layout(gtx)
						}))
				},
			},
			//Cell{
			//	headCell: "Region",
			//	size:     175,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "Subregion",
			//	size:     220,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "International Direct Dial Root",
			//	size:     280,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "Independent",
			//	size:     180,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "Status",
			//	size:     175,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "United Nations Member",
			//	size:     200,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "CCA 2",
			//	size:     85,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "CCA 3",
			//	size:     85,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "CCN 3",
			//	size:     85,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "Area",
			//	size:     100,
			//	Layout:   nil,
			//},
			//Cell{
			//	headCell: "Population",
			//	size:     100,
			//	Layout:   nil,
			//},
		)
	}
}

//row := layout.Flex{}.Layout(gtx,
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(450, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						return material.Body1(th, r.Name).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(550, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						return material.Body1(th, r.OfficialName).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(200, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						Capital := "N/A"
//						if len(r.Capital) > 0 {
//							Capital = r.Capital[0]
//						}
//						return material.Body1(th, Capital).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(175, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						return material.Body1(th, r.Region).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(220, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						return material.Body1(th, r.Subregion).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(280, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						return material.Body1(th, r.IddRoot).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(180, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						Independent := "No"
//						if r.Independent {
//							Independent = "Yes"
//						}
//						return material.Body1(th, Independent).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(175, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						return material.Body1(th, r.Status).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(200, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(func(gtx C) D {
//						UNMember := "No"
//						if r.UNMember {
//							UNMember = "Yes"
//						}
//						return material.Body1(th, UNMember).Layout(gtx)
//					}))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(85, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(material.Body1(th, r.Cca2).Layout))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(85, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(material.Body1(th, r.Cca3).Layout))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(85, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(material.Body1(th, r.Ccn3).Layout))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(100, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(material.Body1(th, fmt.Sprintf("%.0f", r.Area)).Layout))
//			})
//		})
//	}),
//	layout.Rigid(func(gtx C) D {
//		return border.Layout(gtx, func(gtx C) D {
//			return inset.Layout(gtx, func(gtx C) D {
//				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
//
//					layout.Expanded(func(gtx C) D {
//						return g.ColoredArea(gtx, image.Pt(100, gtx.Constraints.Min.Y), rowColor)
//					}),
//
//					layout.Stacked(material.Body1(th, fmt.Sprintf("%d", int(r.Population))).Layout))
//			})
//		})
//	}))
//return row
