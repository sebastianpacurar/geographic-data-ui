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

		columns []cell
	}

	cell struct {
		headCell string
		position int
		sizeX    int
		content  interface{}
	}
)

// generateColumns - holds the state of every column cell of every Row
func (r *Row) generateColumns() {
	r.columns = make([]cell, 0, len(ColNames))
	if len(r.columns) == 0 {
		for i := range ColNames {
			r.columns = append(r.columns,
				cell{headCell: OFFICIAL_NAME, position: i, content: r.OfficialName, sizeX: 550},
				cell{headCell: CAPITALS, position: i, content: r.Capitals, sizeX: 250},
				cell{headCell: REGION, position: i, content: r.Region, sizeX: 175},
				cell{headCell: SUBREGION, position: i, content: r.Subregion, sizeX: 225},
				cell{headCell: CONTINENTS, position: i, content: r.Continents, sizeX: 175},
				cell{headCell: LANGUAGES, position: i, content: r.Languages, sizeX: 650},
				cell{headCell: IDD_ROOT, position: i, content: r.IddRoot, sizeX: 165},
				cell{headCell: IDD_SUFFIXES, position: i, content: r.IddSuffixes, sizeX: 200},
				cell{headCell: TOP_LEVEL_DOMAINS, position: i, content: r.TopLevelDomains, sizeX: 200},
				cell{headCell: INDEPENDENT, position: i, content: r.Independent, sizeX: 180},
				cell{headCell: STATUS, position: i, content: r.Status, sizeX: 175},
				cell{headCell: UNITED_NATIONS_MEMBER, position: i, content: r.UNMember, sizeX: 200},
				cell{headCell: LANDLOCKED, position: i, content: r.Landlocked, sizeX: 180},
				cell{headCell: CCA2, position: i, content: r.Cca2, sizeX: 85},
				cell{headCell: CCA3, position: i, content: r.Cca3, sizeX: 85},
				cell{headCell: CCN3, position: i, content: r.Ccn3, sizeX: 85},
				cell{headCell: CIOC, position: i, content: r.Cioc, sizeX: 95},
				cell{headCell: FIFA, position: i, content: r.Fifa, sizeX: 95},
				cell{headCell: AREA, position: i, content: r.Area, sizeX: 125},
				cell{headCell: POPULATION, position: i, content: r.Population, sizeX: 150},
				cell{headCell: LATITUDE, position: i, content: r.Latitude, sizeX: 150},
				cell{headCell: LONGITUDE, position: i, content: r.Longitude, sizeX: 150},
				cell{headCell: START_OF_WEEK, position: i, content: r.StartOfWeek, sizeX: 150},
				cell{headCell: CAR_SIGNS, position: i, content: r.CarSigns, sizeX: 150},
				cell{headCell: CAR_SIDE, position: i, content: r.CarSide, sizeX: 100})
		}
	}
}

func (r *Row) sortColsBasedOnPos() {
	sort.Slice(r.columns, func(i, j int) bool {
		return r.columns[i].position < r.columns[j].position
	})
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
					if ColState[r.columns[i].headCell] {
						dim = layout.Stack{Alignment: layout.Center}.Layout(gtx,
							layout.Expanded(func(gtx C) D {
								return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(r.columns[i].sizeX))), r.sizeY), rowColor)
							}),
							layout.Stacked(material.Body1(th, r.parseCellContent(r.columns[i].headCell, r.columns[i].content)).Layout))
					}
					return dim
				})
			})
		} else {
			return r.colList.Layout(gtx, len(ColNames), func(gtx C, i int) D {
				var dim D
				if ColState[r.columns[i].headCell] {
					if SearchBy == r.columns[i].headCell {
						rowColor = globals.Colours[colours.LIGHT_YELLOW]
					} else {
						rowColor = globals.Colours[colours.LAVENDERBLUSH]
					}
					dim = layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							return globals.ColoredArea(gtx, image.Pt(gtx.Px(unit.Dp(float32(r.columns[i].sizeX))), r.headerSizeY), rowColor)
						}),
						layout.Stacked(material.Body1(th, r.columns[i].headCell).Layout))
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
				return layout.Stack{Alignment: layout.W}.Layout(gtx,
					layout.Expanded(func(gtx C) D {

						// maintain header row at the same size on cross-axis, no matter the resize boundaries
						r.headerSizeY = gtx.Constraints.Min.Y + gtx.Px(unit.Dp(float32(25)))
						return globals.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, r.headerSizeY), globals.Colours[colours.ELECTRIC_BLUE])
					}),
					layout.Stacked(func(gtx C) D {
						return layout.Inset{Left: unit.Dp(5)}.Layout(gtx, func(gtx C) D {
							return material.Body1(th, fmt.Sprintf("Countries (%d)", GetDisplayedCount())).Layout(gtx)
						})
					}))
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
				return layout.Stack{Alignment: layout.W}.Layout(gtx,
					layout.Expanded(func(gtx C) D {

						// resize entire row based on sticky column cross-axis row size
						r.sizeY = gtx.Constraints.Min.Y + gtx.Px(unit.Dp(float32(20)))
						return globals.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, r.sizeY), cellColor)
					}),
					layout.Stacked(func(gtx C) D {
						return layout.Inset{Left: unit.Dp(5)}.Layout(gtx, func(gtx C) D {
							return material.Body1(th, r.Name).Layout(gtx)
						})
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
