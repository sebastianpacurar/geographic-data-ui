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
	"strconv"
	"strings"
)

var ColNames = []string{
	"Name", "Official Name", "Capital", "Region", "Subregion", "Continents", "IDD Root", "IDD Suffixes", "Top Level Domains",
	"Independent", "Status", "United Nations Member", "Land Locked", "CCA 2", "CCA 3", "CCN 3", "IOC Code", "FIFA Code",
	"Area", "Population", "Latitude", "Longitude", "Start of Week", "Car Signs", "Car Side",
}

type (
	Row struct {
		Name            string
		OfficialName    string
		Capital         []string
		Region          string
		Subregion       string
		Continents      []string
		Languages       map[string]string
		IddRoot         string
		IddSuffixes     []string
		TopLevelDomains []string
		Independent     bool
		Status          string
		UNMember        bool
		Landlocked      bool
		Cca2            string
		Cca3            string
		Ccn3            string
		Cioc            string
		Fifa            string
		Area            float64
		Population      int32
		Latitude        float64
		Longitude       float64
		StartOfWeek     string
		CarSigns        []string
		CarSide         string

		Active     bool
		Selected   bool
		IsCPViewed bool

		Click   widget.Clickable
		colList layout.List
		loaded  bool
		Columns []Cell
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
				return r.colList.Layout(gtx, len(ColNames), func(gtx C, i int) D {
					return r.Columns[i].Layout(gtx, th, &r.Columns[i], rowColor, isHeader)
				})
			})
		} else {
			rowColor = g.Colours[colours.LAVENDERBLUSH]
			return r.colList.Layout(gtx, len(ColNames), func(gtx C, i int) D {
				return r.Columns[i].Layout(gtx, th, &r.Columns[i], rowColor, isHeader)
			})
		}
	})
}

func (r *Row) generateColumns() {
	for range ColNames {
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
							capital := "-"
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
				Size:     225,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							subregion := "-"
							if len(r.Subregion) > 0 {
								subregion = r.Subregion
							}
							if isHeader {
								subregion = c.HeadCell
							}
							return material.Body1(th, subregion).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Continents",
				Size:     175,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := strings.Join(r.Continents, ", ")
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},

			// TODO: fix shallow copy issue on rerendering
			//Cell{
			//	HeadCell: "Languages",
			//	Size:     350,
			//	Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
			//		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			//			layout.Expanded(func(gtx C) D {
			//				return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
			//			}),
			//			layout.Stacked(func(gtx C) D {
			//				langs := make([]string, 0)
			//				for _, v := range r.Languages {
			//					langs = append(langs, v)
			//				}
			//				res := strings.Join(langs, ", ")
			//				if isHeader {
			//					res = c.HeadCell
			//				}
			//				return material.Body1(th, res).Layout(gtx)
			//			}))
			//	},
			//},

			Cell{
				HeadCell: "IDD Root",
				Size:     165,
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
				HeadCell: "IDD Suffixes",
				Size:     200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := ""
							if r.Name == "United States" {
								limits := []string{r.IddSuffixes[0]}
								limits = append(limits, r.IddSuffixes[len(r.IddSuffixes)-1])
								res = strings.Join(limits, "-")
							} else {
								res = strings.Join(r.IddSuffixes, ", ")
							}
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Top Level Domains",
				Size:     200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := "-"
							if len(r.TopLevelDomains) > 0 || r.TopLevelDomains != nil {
								res = strings.Join(r.TopLevelDomains, ", ")
							}
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
				HeadCell: "Land Locked",
				Size:     180,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							landLocked := "No"
							if r.UNMember {
								landLocked = "Yes"
							}
							if isHeader {
								landLocked = c.HeadCell
							}
							return material.Body1(th, landLocked).Layout(gtx)
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
							ccn := "-"
							if len(r.Ccn3) > 0 {
								ccn = r.Ccn3
							}
							if isHeader {
								ccn = c.HeadCell
							}
							return material.Body1(th, ccn).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "IOC Code",
				Size:     95,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							ioc := "-"
							if len(r.Cioc) > 0 {
								ioc = r.Cioc
							}
							if isHeader {
								ioc = c.HeadCell
							}
							return material.Body1(th, ioc).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "FIFA Code",
				Size:     95,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							fifa := "-"
							if len(r.Cioc) > 0 {
								fifa = r.Cioc
							}
							if isHeader {
								fifa = c.HeadCell
							}
							return material.Body1(th, fifa).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Area",
				Size:     125,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := strconv.FormatFloat(r.Area, 'f', -1, 32)
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Population",
				Size:     150,
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
			},
			Cell{
				HeadCell: "Latitude",
				Size:     150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := strconv.FormatFloat(r.Latitude, 'f', -1, 64)
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Longitude",
				Size:     150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := strconv.FormatFloat(r.Longitude, 'f', -1, 64)
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Start of Week",
				Size:     150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.StartOfWeek
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Car Signs",
				Size:     150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := strings.Join(r.CarSigns, ", ")
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Car Side",
				Size:     100,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return g.ColoredArea(gtx, image.Pt(c.Size, gtx.Constraints.Min.Y), color)
						}),
						layout.Stacked(func(gtx C) D {
							res := r.CarSide
							if isHeader {
								res = c.HeadCell
							}
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			})
	}
}
