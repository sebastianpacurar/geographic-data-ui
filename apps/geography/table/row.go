package table

import (
	"fmt"
	"gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var ColNames = []string{
	"Official Name", "Capital", "Region", "Subregion", "Languages", "Continents", "IDD Root", "IDD Suffixes", "Top Level Domains",
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

		Active          bool
		ActiveContinent bool
		Selected        bool
		IsCPViewed      bool

		btn         widget.Clickable
		colList     layout.List
		sizeY       int
		headerSizeY int
		loaded      bool

		columns []Cell
	}

	Cell struct {
		HeadCell  string
		sizeX     int
		IsEnabled bool
		Layout    func(C, *material.Theme, *Cell, color.NRGBA, bool) D
	}
)

func (r *Row) LayRow(gtx C, th *material.Theme, isHeader bool) D {
	rowColor := globals.Colours[colours.ANTIQUE_WHITE]

	if !r.loaded {
		r.colList.Alignment = layout.Middle
		r.generateColumns()
		r.loaded = true
	}

	border := widget.Border{
		Color: globals.Colours[colours.GREY],
		Width: unit.Dp(1),
	}
	if isHeader {
		border.Width = unit.Dp(1.5)
		border.Color = globals.Colours[colours.GREY]
	}

	return border.Layout(gtx, func(gtx C) D {
		if !isHeader {
			return material.Clickable(gtx, &r.btn, func(gtx C) D {
				if r.Selected {
					rowColor = globals.Colours[colours.AERO_BLUE]
				}
				if r.btn.Hovered() {
					if r.Selected {
						rowColor = globals.Colours[colours.LIGHT_SALMON]
					} else {
						rowColor = globals.Colours[colours.NYANZA]
					}
				}
				return r.colList.Layout(gtx, len(ColNames), func(gtx C, i int) D {
					return r.columns[i].Layout(gtx, th, &r.columns[i], rowColor, isHeader)
				})
			})
		} else {
			rowColor = globals.Colours[colours.LAVENDERBLUSH]
			return r.colList.Layout(gtx, len(ColNames), func(gtx C, i int) D {
				return r.columns[i].Layout(gtx, th, &r.columns[i], rowColor, isHeader)
			})
		}
	})
}

// LayNameColumn - Lay sticky country name Column - TODO: simplify!
func (r *Row) LayNameColumn(gtx C, th *material.Theme, isHeader bool) D {
	cellColor := globals.Colours[colours.ANTIQUE_WHITE]
	if !r.loaded {
		r.generateColumns()
		r.loaded = true
	}

	border := widget.Border{
		Color: globals.Colours[colours.GREY],
		Width: unit.Dp(1),
	}

	return border.Layout(gtx, func(gtx C) D {
		if isHeader {
			var btn widget.Clickable
			return material.Clickable(gtx, &btn, func(gtx C) D {
				gtx.Queue = nil
				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx C) D {

						// maintain header row at the same size on cross axis, no matter the resize boundaries
						r.headerSizeY = gtx.Constraints.Min.Y
						return globals.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y), globals.Colours[colours.ELECTRIC_BLUE])
					}),
					layout.Stacked(func(gtx C) D {
						return material.Body1(th, "Country").Layout(gtx)
					}))
			})
		} else {
			return material.Clickable(gtx, &r.btn, func(gtx C) D {
				if r.Selected {
					cellColor = globals.Colours[colours.AERO_BLUE]
				}
				if r.btn.Hovered() {
					if r.Selected {
						cellColor = globals.Colours[colours.LIGHT_SALMON]
					} else {
						cellColor = globals.Colours[colours.NYANZA]
					}
				}
				return layout.Stack{}.Layout(gtx,
					layout.Expanded(func(gtx C) D {

						// resize entire row based on sticky column cross axis size
						r.sizeY = gtx.Constraints.Min.Y
						return globals.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y), cellColor)
					}),
					layout.Stacked(func(gtx C) D {
						return material.Body1(th, r.Name).Layout(gtx)
					}))
			})
		}

	})
}

func (r *Row) generateColumns() {
	r.columns = make([]Cell, 0, len(ColNames))
	for _ = range ColNames {
		r.columns = append(r.columns,
			//Cell{
			//	HeadCell: "Name",
			//	sizeX:     450,
			//	Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
			//		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			//			layout.Expanded(func(gtx C) D {
			//				return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), r.sizeY), color)
			//			}),
			//			layout.Stacked(func(gtx C) D {
			//				res := r.Name
			//				if isHeader {
			//					res = c.HeadCell
			//				}
			//				return material.Body1(th, res).Layout(gtx)
			//			}))
			//	},
			//},
			Cell{
				HeadCell: "Official Name",
				sizeX:    550,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := r.OfficialName
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Capital",
				sizeX:    200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					capital := "-"
					sizeCross := r.sizeY
					if len(r.Capital) > 0 {
						capital = r.Capital[0]
					}
					if isHeader {
						capital = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, capital).Layout))
				},
			},
			Cell{
				HeadCell: "Region",
				sizeX:    175,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := r.Region
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Subregion",
				sizeX:    225,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					subregion := "-"
					sizeCross := r.sizeY
					if len(r.Subregion) > 0 {
						subregion = r.Subregion
					}
					if isHeader {
						subregion = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, subregion).Layout))
				},
			},
			Cell{
				HeadCell: "Continents",
				sizeX:    175,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := strings.Join(r.Continents, ", ")
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Languages",
				sizeX:    650,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := ""
					sizeCross := r.sizeY
					langs := make([]string, 0, len(r.Languages))

					for _, v := range r.Languages {
						langs = append(langs, v)
					}
					sort.Strings(langs)
					if len(langs) <= 5 {
						res = strings.Join(langs, ", ")
					} else {

						// first 5 + (all - 5) more
						res = strings.Join(langs[:5], ", ")
						res += fmt.Sprintf(" + %d more", len(langs[5:]))
					}
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "IDD Root",
				sizeX:    165,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := r.IddRoot
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "IDD Suffixes",
				sizeX:    200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := ""
					sizeCross := r.sizeY
					if r.Name == "United States" {
						limits := []string{r.IddSuffixes[0]}
						limits = append(limits, r.IddSuffixes[len(r.IddSuffixes)-1])
						res = strings.Join(limits, "-")
					} else {
						res = strings.Join(r.IddSuffixes, ", ")
					}
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Top Level Domains",
				sizeX:    200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := "-"
					sizeCross := r.sizeY
					if len(r.TopLevelDomains) > 0 || r.TopLevelDomains != nil {

						// exclude non latin characters, for now
						latinList := make([]string, 0, len(r.TopLevelDomains))
						for i := range r.TopLevelDomains {
							isLatin := true
							for _, v := range r.TopLevelDomains[i][1:] {
								if !unicode.In(v, unicode.Latin) {
									isLatin = false
									break
								}
							}
							if isLatin {
								latinList = append(latinList, r.TopLevelDomains[i])
							}
						}

						if len(latinList) >= 0 {
							res = strings.Join(latinList, ", ")
						}
					}
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Independent",
				sizeX:    180,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					independent := "No"
					sizeCross := r.sizeY
					if r.Independent {
						independent = "Yes"
					}
					if isHeader {
						independent = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, independent).Layout))
				},
			},
			Cell{
				HeadCell: "Status",
				sizeX:    175,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := r.Status
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "United Nations Member",
				sizeX:    200,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					unMember := "No"
					sizeCross := r.sizeY
					if r.UNMember {
						unMember = "Yes"
						sizeCross = r.headerSizeY
					}
					if isHeader {
						unMember = c.HeadCell
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, unMember).Layout))
				},
			},
			Cell{
				HeadCell: "Land Locked",
				sizeX:    180,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					landLocked := "No"
					sizeCross := r.sizeY
					if r.Landlocked {
						landLocked = "Yes"
					}
					if isHeader {
						landLocked = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, landLocked).Layout))
				},
			},
			Cell{
				HeadCell: "CCA 2",
				sizeX:    85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := r.Cca2
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "CCA 3",
				sizeX:    85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := r.Cca3
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "CCN 3",
				sizeX:    85,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					ccn := "-"
					sizeCross := r.sizeY
					if len(r.Ccn3) > 0 {
						ccn = r.Ccn3
					}
					if isHeader {
						ccn = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, ccn).Layout))
				},
			},
			Cell{
				HeadCell: "IOC Code",
				sizeX:    95,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					ioc := "-"
					sizeCross := r.sizeY
					if len(r.Cioc) > 0 {
						ioc = r.Cioc
					}
					if isHeader {
						ioc = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, ioc).Layout))
				},
			},
			Cell{
				HeadCell: "FIFA Code",
				sizeX:    95,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					fifa := "-"
					sizeCross := r.sizeY
					if len(r.Fifa) > 0 {
						fifa = r.Fifa
					}
					if isHeader {
						fifa = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, fifa).Layout))
				},
			},
			Cell{
				HeadCell: "Area",
				sizeX:    125,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := strconv.FormatFloat(r.Area, 'f', -1, 32)
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Population",
				sizeX:    150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := fmt.Sprintf("%d", int(r.Population))
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Latitude",
				sizeX:    150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := strconv.FormatFloat(r.Latitude, 'f', -1, 64)
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Longitude",
				sizeX:    150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := strconv.FormatFloat(r.Longitude, 'f', -1, 64)
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(func(gtx C) D {
							return material.Body1(th, res).Layout(gtx)
						}))
				},
			},
			Cell{
				HeadCell: "Start of Week",
				sizeX:    150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := r.StartOfWeek
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Car Signs",
				sizeX:    150,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := "-"
					sizeCross := r.sizeY
					if len(r.CarSigns) > 0 || r.CarSigns != nil {
						res = strings.Join(r.CarSigns, ", ")
					}
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			},
			Cell{
				HeadCell: "Car Side",
				sizeX:    100,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := r.CarSide
					sizeCross := r.sizeY
					if isHeader {
						res = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, res).Layout))
				},
			})
	}
}
