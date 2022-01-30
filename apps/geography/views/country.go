package views

import (
	"fmt"
	"gioui-experiment/apps/geography/data"
	"gioui-experiment/apps/geography/table"
	"gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const MILE_VAL = 1.609344

var (
	// contains Area vals in this format [Miles, Kilometers]
	areaValues []float64
	areas      = []string{"Miles", "Kilometers"}
)

type (
	C = layout.Context
	D = layout.Dimensions

	CountryView struct {
		leftList  layout.List
		rightList layout.List

		txtWraps []txtWrap
		loaded   bool
	}

	txtWrap struct {
		prop    string
		content interface{}
	}
)

func (cv *CountryView) Layout(gtx C, th *material.Theme, country data.Country) D {
	cv.leftList.Axis = layout.Vertical
	areaValues = make([]float64, 2)
	areaValues[0] = country.Area
	areaValues[1] = country.Area * MILE_VAL
	cv.generateTxtWrap(country)

	return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx C) D {
		return layout.Flex{WeightSum: 2}.Layout(gtx,
			layout.Flexed(1, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return cv.leftList.Layout(gtx, len(cv.txtWraps), func(gtx C, i int) D {
							return cv.txtWraps[i].LayTextWrap(gtx, th, country)
						})
					}))
			}),
			layout.Flexed(1, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Flexed(1, func(gtx C) D {
						return widget.Image{
							Src: paint.NewImageOp(country.FlagImg),
							Fit: widget.Contain,
						}.Layout(gtx)
					}))
			}))
	})
}

// generateTxtWrap - Generates the div+label wrapper
func (cv *CountryView) generateTxtWrap(country data.Country) {
	cv.txtWraps = make([]txtWrap, 0, 4)
	cv.txtWraps = append(cv.txtWraps,
		txtWrap{prop: table.OFFICIAL_NAME, content: country.Name.Official},
		txtWrap{prop: table.CAPITALS, content: country.Capitals},
		txtWrap{prop: table.POPULATION, content: country.Population},
		txtWrap{prop: table.AREA, content: country.Area},
		txtWrap{prop: table.REGION, content: country.Region})
}

// LayTextWrap - Lays the data with a subheading divider and label(s)
func (tw *txtWrap) LayTextWrap(gtx C, th *material.Theme, country data.Country) D {
	var (
		lbl material.LabelStyle
		div component.DividerStyle
	)

	// render a special version of labels for AREA
	if tw.prop != table.AREA {
		lbl = material.Body2(th, getValue(country, tw.prop, tw.content))
		lbl.TextSize = th.TextSize.Scale(15.0 / 16.0)
	}

	div = component.SubheadingDivider(th, tw.prop)
	div.Subheading = material.Body2(th, tw.prop)

	div.Right = unit.Dp(15)
	div.Fill = globals.Colours[colours.LIGHT_SEA_GREEN]
	div.Subheading.TextSize = th.TextSize.Scale(15.0 / 16.0)
	div.Subheading.Font.Weight = text.SemiBold
	div.Subheading.Color = globals.Colours[colours.SEA_GREEN]

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(div.Layout),
		layout.Rigid(func(gtx C) D {
			return layout.Inset{Left: unit.Dp(15)}.Layout(gtx, func(gtx C) D {
				var dims D

				// Format Area Layout to contain Miles and Kilometers values
				if tw.prop == table.AREA {
					dims = layout.Flex{}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							areaList := layout.List{Axis: layout.Vertical}
							return areaList.Layout(gtx, 2, func(gtx C, i int) D {
								var (
									areaLblVal  material.LabelStyle
									areaLblType material.LabelStyle
								)
								return layout.Flex{}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										// set min-width on Main Axis to keep areaLblType labels aligned on Cross Axis
										gtx.Constraints.Min.X = gtx.Px(unit.Dp(float32(100)))
										areaLblVal = material.Body2(th, getValue(country, tw.prop, areaValues[i]))
										return areaLblVal.Layout(gtx)

									}),
									layout.Rigid(func(gtx C) D {
										areaLblType = material.Body2(th, getValue(country, tw.prop, areas[i]))
										return areaLblType.Layout(gtx)
									}))
							})
						}),
					)
				} else {
					dims = lbl.Layout(gtx)
				}
				return dims
			})
		}),
	)
}

// getValue - returns the content interface value
func getValue(country data.Country, property string, content interface{}) string {
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
			switch property {
			case table.TOP_LEVEL_DOMAINS:
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

			case table.IDD_SUFFIXES:
				if country.Name.Common == "United States" {
					limits := []string{country.Idd.Suffixes[0]}
					limits = append(limits, country.Idd.Suffixes[len(country.Idd.Suffixes)-1])
					res = strings.Join(limits, "-")
				} else {
					res = strings.Join(country.Idd.Suffixes, ", ")
				}
			default:
				res = strings.Join(parsed, ", ")
			}
		}

	case reflect.Map:
		kvPair := reflect.ValueOf(content)
		switch property {
		case table.LANGUAGES:
			if country.Name.Common == "Antarctica" {
				res = "-"
			} else {
				parsed := make([]string, 0, len(country.Languages))
				for _, el := range kvPair.MapKeys() {
					parsed = append(parsed, kvPair.MapIndex(el).String())
				}
				sort.Strings(parsed)
				res = strings.Join(parsed, ", ")
			}
		}
	}
	return res
}
