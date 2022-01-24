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
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type (
	Row struct {
		Name            string
		OfficialName    string
		Capitals        []string
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

		IsSearchedFor     bool
		IsActiveContinent bool
		Selected          bool

		btn         widget.Clickable
		colList     layout.List
		sizeY       int
		headerSizeY int
		loaded      bool

		Columns []cell
	}

	cell struct {
		HeadCell string
		sizeX    int
		content  interface{}
	}
)

// generateColumns - holds the state of every column cell of every Row
func (r *Row) generateColumns() {
	r.Columns = make([]cell, 0, len(ColNames))
	for range ColNames {
		r.Columns = append(r.Columns,
			cell{HeadCell: OFFICIAL_NAME, content: r.OfficialName, sizeX: 550},
			cell{HeadCell: CAPITALS, content: r.Capitals, sizeX: 250},
			cell{HeadCell: REGION, content: r.Region, sizeX: 175},
			cell{HeadCell: SUBREGION, content: r.Subregion, sizeX: 225},
			cell{HeadCell: CONTINENTS, content: r.Continents, sizeX: 175},
			cell{HeadCell: LANGUAGES, content: r.Languages, sizeX: 650},
			cell{HeadCell: IDD_ROOT, content: r.IddRoot, sizeX: 165},
			cell{HeadCell: IDD_SUFFIXES, content: r.IddSuffixes, sizeX: 200},
			cell{HeadCell: TOP_LEVEL_DOMAINS, content: r.TopLevelDomains, sizeX: 200},
			cell{HeadCell: INDEPENDENT, content: r.Independent, sizeX: 180},
			cell{HeadCell: STATUS, content: r.Status, sizeX: 175},
			cell{HeadCell: UNITED_NATIONS_MEMBER, content: r.UNMember, sizeX: 200},
			cell{HeadCell: LANDLOCKED, content: r.Landlocked, sizeX: 180},
			cell{HeadCell: CCA2, content: r.Cca2, sizeX: 85},
			cell{HeadCell: CCA3, content: r.Cca3, sizeX: 85},
			cell{HeadCell: CCN3, content: r.Ccn3, sizeX: 85},
			cell{HeadCell: CIOC, content: r.Cioc, sizeX: 95},
			cell{HeadCell: FIFA, content: r.Fifa, sizeX: 95},
			cell{HeadCell: AREA, content: r.Area, sizeX: 125},
			cell{HeadCell: POPULATION, content: r.Population, sizeX: 150},
			cell{HeadCell: LATITUDE, content: r.Latitude, sizeX: 150},
			cell{HeadCell: LONGITUDE, content: r.Longitude, sizeX: 150},
			cell{HeadCell: START_OF_WEEK, content: r.StartOfWeek, sizeX: 150},
			cell{HeadCell: CAR_SIGNS, content: r.CarSigns, sizeX: 150},
			cell{HeadCell: CAR_SIDE, content: r.CarSide, sizeX: 100})
	}
}

// parseCellContent - stringify the content country cell data
func (r *Row) parseCellContent(headCell string, content interface{}) string {
	res := "-"
	if content == nil {
		return res
	}
	switch reflect.TypeOf(content).Kind() {
	case reflect.String:
		if len(reflect.ValueOf(content).String()) > 0 {
			res = reflect.ValueOf(content).String()
		}
	case reflect.Float64:
		res = strconv.FormatFloat(reflect.ValueOf(content).Float(), 'f', -1, 32)
	case reflect.Int32:
		res = fmt.Sprintf("%d", int(reflect.ValueOf(content).Int()))
	case reflect.Bool:
		res = "No"
		if reflect.ValueOf(content).Bool() {
			res = "Yes"
		}

	case reflect.Slice:
		arr := reflect.ValueOf(content)
		if arr.Len() > 0 {
			var parsed []string
			for i := 0; i < arr.Len(); i++ {
				if arr.Index(i).String() == "" {
					return res
				}
				parsed = append(parsed, arr.Index(i).String())
			}
			switch headCell {
			case TOP_LEVEL_DOMAINS:
				// exclude non latin characters, for now
				filtered := make([]string, 0, len(parsed))
				for i := range parsed {
					isLatin := true
					for _, v := range parsed[i][1:] {
						if !unicode.In(v, unicode.Latin) {
							isLatin = false
							break
						}
					}
					if isLatin {
						filtered = append(filtered, parsed[i])
					}
				}
				res = strings.Join(filtered, ", ")

			case IDD_SUFFIXES:
				if r.Name == "United States" {
					limits := []string{r.IddSuffixes[0]}
					limits = append(limits, r.IddSuffixes[len(r.IddSuffixes)-1])
					res = strings.Join(limits, "-")
				} else {
					res = strings.Join(r.IddSuffixes, ", ")
				}
			default:
				res = strings.Join(parsed, ", ")
			}
		}

	case reflect.Map:
		kvPair := reflect.ValueOf(content)
		switch headCell {
		case LANGUAGES:
			parsed := make([]string, 0, len(r.Languages))
			for _, el := range kvPair.MapKeys() {
				parsed = append(parsed, kvPair.MapIndex(el).String())
			}
			sort.Strings(parsed)
			if len(parsed) <= 5 {
				res = strings.Join(parsed, ", ")
			} else {

				// first 5 + (all - 5) more
				res = strings.Join(parsed[:5], ", ")
				res += fmt.Sprintf(" + %d more", len(parsed[5:]))
			}
		}
	}
	return res
}

// LayRow - Lay the row with all column cells parsed
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
					var dim D
					if ColState[r.Columns[i].HeadCell] {
						dim = layout.Stack{Alignment: layout.Center}.Layout(gtx,
							layout.Expanded(func(gtx C) D {
								return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(r.Columns[i].sizeX))), r.sizeY), rowColor)
							}),
							layout.Stacked(material.Body1(th, r.parseCellContent(r.Columns[i].HeadCell, r.Columns[i].content)).Layout))
					}
					return dim
				})
			})
		} else {
			return r.colList.Layout(gtx, len(ColNames), func(gtx C, i int) D {
				var dim D
				if ColState[r.Columns[i].HeadCell] {
					if SearchBy == r.Columns[i].HeadCell {
						rowColor = globals.Colours[colours.LIGHT_YELLOW]
					} else {
						rowColor = globals.Colours[colours.LAVENDERBLUSH]
					}
					dim = layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(r.Columns[i].sizeX))), r.headerSizeY), rowColor)
						}),
						layout.Stacked(material.Body1(th, r.Columns[i].HeadCell).Layout))
				}
				return dim
			})
		}
	})
}

// LayNameColumn - Lay the sticky Name.Common column
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
					layout.Stacked(material.Body1(th, fmt.Sprintf("Country (%d)", GetDisplayedCount())).Layout))
			})
		} else {
			return material.Clickable(gtx, &r.btn, func(gtx C) D {
				if SearchBy == "Name" {
					cellColor = globals.Colours[colours.LIGHT_YELLOW]
				}
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

// GetDisplayedCount - returns the number of displayed countries as rows or cards
func GetDisplayedCount() int {
	count := 0
	for i := range data.Cached {
		if data.Cached[i].IsSearchedFor && data.Cached[i].IsActiveContinent {
			count++
		}
	}
	return count
}
