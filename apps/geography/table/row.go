package table

import (
	"fmt"
	"gioui-experiment/apps/geography/data"
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

var (
	ColNames = []string{
		OFFICIAL_NAME, CAPITAL, REGION, SUBREGION, LANGUAGES, CONTINENTS, IDD_ROOT, IDD_SUFFIXES, TOP_LEVEL_DOMAINS,
		INDEPENDENT, STATUS, UNITED_NATIONS_MEMBER, LANDLOCKED, CCA2, CCA3, CCN3, CIOC, FIFA, AREA, POPULATION, LATITUDE,
		LONGITUDE, START_OF_WEEK, CAR_SIGNS, CAR_SIDE,
	}
)

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

		btn         widget.Clickable
		colList     layout.List
		sizeY       int
		headerSizeY int
		loaded      bool

		Columns []Cell
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
		r.GenerateColumns()
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
					var dim D
					if ColsState[r.Columns[i].HeadCell] {
						dim = r.Columns[i].Layout(gtx, th, &r.Columns[i], rowColor, isHeader)
					}
					return dim
				})
			})
		} else {
			rowColor = globals.Colours[colours.LAVENDERBLUSH]
			return r.colList.Layout(gtx, len(ColNames), func(gtx C, i int) D {
				var dim D
				if ColsState[r.Columns[i].HeadCell] {
					dim = r.Columns[i].Layout(gtx, th, &r.Columns[i], rowColor, isHeader)
				}
				return dim
			})
		}
	})
}

// LayNameColumn - Lay sticky country name Column - TODO: simplify!
func (r *Row) LayNameColumn(gtx C, th *material.Theme, isHeader bool) D {
	cellColor := globals.Colours[colours.ANTIQUE_WHITE]

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

						// maintain header row at the same size on cross-axis, no matter the resize boundaries
						r.headerSizeY = gtx.Constraints.Min.Y
						return globals.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y), globals.Colours[colours.ELECTRIC_BLUE])
					}),
					layout.Stacked(material.Body1(th, fmt.Sprintf("Country (%d)", data.GetDisplayedCount())).Layout))
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

						// resize entire row based on sticky column cross-axis row size
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

func (r *Row) GenerateColumns() {
	r.Columns = make([]Cell, 0, len(ColNames))
	for range ColNames {
		r.Columns = append(r.Columns,
			Cell{
				HeadCell:  OFFICIAL_NAME,
				sizeX:     550,
				IsEnabled: true,
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
				HeadCell:  CAPITAL,
				sizeX:     200,
				IsEnabled: true,
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
				HeadCell:  REGION,
				sizeX:     175,
				IsEnabled: true,
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
				HeadCell:  SUBREGION,
				sizeX:     225,
				IsEnabled: true,
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
				HeadCell:  CONTINENTS,
				sizeX:     175,
				IsEnabled: true,
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
				HeadCell:  LANGUAGES,
				sizeX:     650,
				IsEnabled: true,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := "-"
					sizeCross := r.sizeY

					if r.Languages != nil {
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
				HeadCell:  IDD_ROOT,
				sizeX:     165,
				IsEnabled: true,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := "-"
					if len(r.IddRoot) > 0 {
						res = r.IddRoot
					}
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
				HeadCell:  IDD_SUFFIXES,
				sizeX:     200,
				IsEnabled: true,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := "-"
					sizeCross := r.sizeY
					if len(r.IddSuffixes) > 0 && len(r.IddSuffixes[0]) > 0 {
						if r.Name == "United States" {
							limits := []string{r.IddSuffixes[0]}
							limits = append(limits, r.IddSuffixes[len(r.IddSuffixes)-1])
							res = strings.Join(limits, "-")
						} else {
							res = strings.Join(r.IddSuffixes, ", ")
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
				HeadCell:  TOP_LEVEL_DOMAINS,
				sizeX:     200,
				IsEnabled: true,
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
				HeadCell:  INDEPENDENT,
				sizeX:     180,
				IsEnabled: true,
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
				HeadCell:  STATUS,
				sizeX:     175,
				IsEnabled: true,
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
				HeadCell:  UNITED_NATIONS_MEMBER,
				sizeX:     200,
				IsEnabled: true,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					unMember := "No"
					sizeCross := r.sizeY
					if r.UNMember {
						unMember = "Yes"
					}
					if isHeader {
						unMember = c.HeadCell
						sizeCross = r.headerSizeY
					}
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(c.sizeX))), sizeCross), color)
						}),
						layout.Stacked(material.Body1(th, unMember).Layout))
				},
			},
			Cell{
				HeadCell:  LANDLOCKED,
				sizeX:     180,
				IsEnabled: true,
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
				HeadCell:  CCA2,
				sizeX:     85,
				IsEnabled: true,
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
				HeadCell:  CCA3,
				sizeX:     85,
				IsEnabled: true,
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
				HeadCell:  CCN3,
				sizeX:     85,
				IsEnabled: true,
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
				HeadCell:  CIOC,
				sizeX:     95,
				IsEnabled: true,
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
				HeadCell:  FIFA,
				sizeX:     95,
				IsEnabled: true,
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
				HeadCell:  AREA,
				sizeX:     125,
				IsEnabled: true,
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
				HeadCell:  POPULATION,
				sizeX:     150,
				IsEnabled: true,
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
				HeadCell:  LATITUDE,
				sizeX:     150,
				IsEnabled: true,
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
				HeadCell:  LONGITUDE,
				sizeX:     150,
				IsEnabled: true,
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
				HeadCell:  START_OF_WEEK,
				sizeX:     150,
				IsEnabled: true,
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
				HeadCell:  CAR_SIGNS,
				sizeX:     150,
				IsEnabled: true,
				Layout: func(gtx C, th *material.Theme, c *Cell, color color.NRGBA, isHeader bool) D {
					res := "-"
					sizeCross := r.sizeY
					if len(r.CarSigns) > 0 && len(r.CarSigns[0]) > 0 {
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
				HeadCell:  CAR_SIDE,
				sizeX:     100,
				IsEnabled: true,
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
