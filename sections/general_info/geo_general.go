package general_info

import (
	"fmt"
	"gioui-experiment/globals"
	"gioui-experiment/sections"
	"gioui-experiment/sections/general_info/data"
	"gioui-experiment/sections/general_info/views/country"
	"gioui-experiment/sections/general_info/views/grid"
	"gioui-experiment/sections/general_info/views/table"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"github.com/xuri/excelize/v2"
	"image"
	"image/color"
	"log"
	"strconv"
	"strings"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Section struct {
		*sections.Router
		th *material.Theme

		ControlPanel
		Display
		cv country.CountryView

		pinBtn       widget.Clickable
		DisableCPBtn widget.Clickable
		isCPDisabled bool
	}

	Display struct {
		// search
		SearchField component.TextField
		CurrentStr  string

		// layout buttons
		TableBtn   widget.Clickable
		GridBtn    widget.Clickable
		SaveAsXlsx widget.Clickable

		// Grid and Table displays
		Grid     grid.Grid
		Table    table.Table
		Selected interface{}

		// mainly used to avoid overhead on "All Cards"
		AllContinents  []string
		Continents     []Continent
		ContinentsList layout.List

		// api data
		Api data.Countries

		// Contextual Viewed Country
		ContextualSet     bool
		ContextualCountry data.Country

		initialSetup bool
		slider       globals.Slider
	}

	Continent struct {
		Name       string
		Btn        widget.Clickable
		IsSelected bool
	}
)

func New(router *sections.Router) *Section {
	return &Section{
		Router: router,
		th:     material.NewTheme(gofont.Collection()),
	}
}

func (s *Section) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Tag: &s.DisableCPBtn,
			},
			Layout: func(gtx C, bg, fg color.NRGBA) D {
				var lbl string
				if s.DisableCPBtn.Clicked() {
					s.isCPDisabled = !s.isCPDisabled
				}
				if !s.isCPDisabled {
					lbl = "Disable CP"
				} else {
					lbl = "Enable CP"
				}
				return material.Button(s.th, &s.DisableCPBtn, lbl).Layout(gtx)
			},
		},
	}
}

func (s *Section) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (s *Section) NavItem() component.NavItem {
	return component.NavItem{
		Name: "General Information",
	}
}

func (s *Section) IsCPDisabled() bool {
	return s.isCPDisabled
}

func (s *Section) LayoutView(gtx C, th *material.Theme) D {
	err := s.Api.InitCountries()
	if err != nil {
		// in case of no internet connection, or if the main request fails
		return material.H6(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	} else {
		s.SearchByColumn(table.SearchBy)

		// run only once at start
		if !s.initialSetup {
			s.SearchField.SingleLine = true

			// initialize all flags
			data.ProcessFlags()

			// initialize Table View at start
			s.Selected = s.Display.Table

			// initialize continents to "All" as Selected
			s.initContinents()
			for i := range s.Continents {
				if s.Continents[i].Name == "All" {
					s.Continents[i].IsSelected = true
					for j := range data.Cached {
						data.Cached[j].IsActiveContinent = true
					}
					break
				}
			}

			// set all countries to IsSearchedFor = true
			for i := range data.Cached {
				data.Cached[i].IsSearchedFor = true
			}
			s.initialSetup = true
		}

		for _, e := range s.AppBar.Events(gtx) {
			switch e.(type) {
			case component.AppBarContextMenuDismissed:
				s.Display.ContextualSet = false
				s.Display.Grid.Contextual = nil
				for i := range data.Cached {
					if data.Cached[i].IsCtxtActive {
						data.Cached[i].IsCtxtActive = false
					}
				}
				op.InvalidateOp{}.Add(gtx.Ops)

			case component.AppBarNavigationClicked:
				s.ModalNavDrawer.Appear(gtx.Now)
				s.NavAnim.Disappear(gtx.Now)
			}
		}

		var dims D

		switch s.Grid.Contextual.(type) {
		case data.Country:
			if !s.ContextualSet {
				s.Router.AppBar.SetContextualActions(
					[]component.AppBarAction{
						component.SimpleIconAction(&s.pinBtn, globals.PinIcon,
							component.OverflowAction{
								Name: "Pin Country",
								Tag:  &s.pinBtn,
							},
						),
					},
					[]component.OverflowAction{},
				)

				for i := range data.Cached {
					if data.Cached[i].IsCtxtActive {
						s.Display.ContextualCountry = data.Cached[i]
					}
				}
				s.Router.AppBar.ToggleContextual(gtx.Now, s.Display.ContextualCountry.Name.Common)
				s.ContextualSet = true
				op.InvalidateOp{}.Add(gtx.Ops)
			}

			// layout Country View
			dims = s.cv.Layout(gtx, th, s.ContextualCountry)

		case nil:
			dims = layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {

					// search field
					search := layout.Flexed(1, func(gtx C) D {
						return s.Display.SearchField.Layout(gtx, th, fmt.Sprintf("Search by %s", table.SearchBy))
					})

					// Table button
					tblBtn := layout.Rigid(func(gtx C) D {
						if s.Display.TableBtn.Clicked() {
							s.Display.Selected = s.Display.Table
							s.Display.slider.PushRight()
						}
						switch s.Display.Selected.(type) {
						case table.Table:
							gtx = gtx.Disabled()
						}
						return layout.Inset{
							Top:    unit.Dp(10),
							Right:  unit.Dp(8),
							Bottom: unit.Dp(8),
							Left:   unit.Dp(8),
						}.Layout(gtx, material.Button(th, &s.Display.TableBtn, "Table").Layout)
					})

					// grid button
					grdBtn := layout.Rigid(func(gtx C) D {
						if s.Display.GridBtn.Clicked() {
							s.Display.Selected = s.Display.Grid
							s.Display.slider.PushLeft()
						}
						switch s.Display.Selected.(type) {
						case grid.Grid:
							gtx = gtx.Disabled()
						}
						return layout.Inset{
							Top:    unit.Dp(10),
							Right:  unit.Dp(8),
							Bottom: unit.Dp(8),
							Left:   unit.Dp(8),
						}.Layout(gtx, material.Button(th, &s.Display.GridBtn, "Grid").Layout)
					})

					//Export to excel button
					exportBtn := layout.Rigid(func(gtx C) D {
						if s.Display.SaveAsXlsx.Clicked() {

						}
						return layout.Inset{
							Top:    unit.Dp(10),
							Right:  unit.Dp(10),
							Bottom: unit.Dp(8),
							Left:   unit.Dp(8),
						}.Layout(gtx, func(gtx C) D {
							var btn material.IconButtonStyle
							btn = material.IconButton(th, &s.Display.SaveAsXlsx, globals.ExcelExportIcon, "Export to Excel")
							btn.Size = unit.Dp(20)
							btn.Inset = layout.UniformInset(unit.Dp(8))
							return btn.Layout(gtx)
						})
					})

					return layout.Inset{Bottom: unit.Dp(15)}.Layout(gtx, func(gtx C) D {
						return layout.Flex{}.Layout(gtx,
							search,
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Alignment: layout.End}.Layout(gtx,
									tblBtn,
									grdBtn,
									exportBtn,
								)
							}))
					})
				}),

				// Select Continent
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Bottom: unit.Dp(15)}.Layout(gtx, func(gtx C) D {
						return layout.Flex{}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.Inset{Top: unit.Dp(10)}.Layout(gtx, func(gtx C) D {
									return material.Body2(th, "Continents: ").Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx C) D {
								return s.ContinentsList.Layout(gtx, len(s.AllContinents), func(gtx C, i int) D {
									var (
										dim D
										btn material.ButtonStyle
									)
									btn = material.Button(th, &s.Continents[i].Btn, s.Continents[i].Name)
									btn.CornerRadius = unit.Dp(1)
									btn.Inset = layout.UniformInset(unit.Dp(10))
									btn.Background = globals.Colours[globals.WHITE]
									btn.Color = globals.Colours[globals.BLACK]
									dim = btn.Layout(gtx)

									if s.Continents[i].Btn.Clicked() {
										name := s.Continents[i].Name
										s.Continents[i].IsSelected = true

										// update the tab ui
										for j := range s.Continents {
											if name != s.Continents[j].Name {
												s.Continents[j].IsSelected = false
											}
										}

										// update IsActiveContinent state
										for j := range data.Cached {
											if name == "All" {
												data.Cached[j].IsActiveContinent = true
											} else {
												continents := strings.Join(data.Cached[j].Continents, " ")
												if strings.Contains(continents, s.Continents[i].Name) {
													data.Cached[j].IsActiveContinent = true
												} else {
													data.Cached[j].IsActiveContinent = false
												}
											}
										}
										op.InvalidateOp{}.Add(gtx.Ops)
									}

									if s.Continents[i].IsSelected {
										dim = widget.Border{
											Width:        unit.Dp(1),
											CornerRadius: btn.CornerRadius,
										}.Layout(gtx, func(gtx C) D {
											size := image.Pt(dim.Size.X, dim.Size.Y)
											gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))

											return layout.Stack{Alignment: layout.S}.Layout(gtx,
												layout.Expanded(func(gtx C) D {
													return globals.ColoredArea(gtx, size, globals.Colours[globals.AERO_BLUE])
												}),
												layout.Stacked(func(gtx C) D {
													return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
														layout.Flexed(1, func(gtx C) D {
															var lbl material.LabelStyle
															lbl = material.Body1(th, s.Continents[i].Name)
															lbl.TextSize = btn.TextSize

															return layout.Flex{}.Layout(gtx,
																layout.Flexed(1, func(gtx C) D {
																	return layout.Center.Layout(gtx, lbl.Layout)
																}))
														}),
														layout.Rigid(func(gtx C) D {
															return layout.Stack{}.Layout(gtx,
																layout.Expanded(func(gtx C) D {
																	return globals.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, 3), globals.Colours[globals.SEA_GREEN])
																}))
														}))
												}))
										})
									}
									return dim
								})
							}),
						)
					})
				}),

				// Selected display
				layout.Rigid(func(gtx C) D {
					return s.Display.slider.Layout(gtx, layout.Horizontal, func(gtx C) D {
						var dims D
						switch s.Display.Selected.(type) {
						case grid.Grid:
							dims = s.Display.Grid.Layout(gtx, th)
						case table.Table:
							dims = s.Display.Table.Layout(gtx, th)
						}
						return dims
					})
				}),
			)
		}
		return dims
	}
}

func (s *Section) LayoutController(gtx C, th *material.Theme) D {
	return s.ControlPanel.Layout(gtx, th)
}

// SearchByColumn - filter countries based on data.Cached and data.Cached manipulation
func (d *Display) SearchByColumn(SearchBy string) {
	if d.CurrentStr != d.SearchField.Text() {
		if d.SearchField.Len() > 0 {
			for i := range data.Cached {
				res := "-"

				// Leaving this as verbose as is, for now
				switch SearchBy {
				case table.NAME:
					res = data.Cached[i].Name.Common
				case table.OFFICIAL_NAME:
					res = data.Cached[i].Name.Official
				case table.CAPITALS:
					if len(data.Cached[i].Capitals) > 0 {
						res = data.Cached[i].Capitals[0]
					}
				case table.TOP_LEVEL_DOMAINS:
					if len(data.Cached[i].TopLevelDomain) > 0 {
						res = data.Cached[i].TopLevelDomain[0]
					}
				case table.INDEPENDENT:
					if data.Cached[i].Independent {
						res = "yes"
					} else {
						res = "no"
					}
				case table.UNITED_NATIONS_MEMBER:
					if data.Cached[i].UNMember {
						res = "yes"
					} else {
						res = "no"
					}
				case table.LANDLOCKED:
					if data.Cached[i].Landlocked {
						res = "yes"
					} else {
						res = "no"
					}
				case table.CCA2:
					res = data.Cached[i].Cca2
				case table.CCA3:
					res = data.Cached[i].Cca3
				case table.CCN3:
					res = data.Cached[i].Ccn3
				case table.CAR_SIDE:
					res = data.Cached[i].Car.Side
				}

				if strings.HasPrefix(strings.ToLower(res), strings.ToLower(d.SearchField.Text())) {
					data.Cached[i].IsSearchedFor = true
				} else {
					data.Cached[i].IsSearchedFor = false
				}
			}
			d.CurrentStr = d.SearchField.Text()
		} else if d.SearchField.Len() == 0 {
			for i := range data.Cached {
				data.Cached[i].IsSearchedFor = true
			}
			d.CurrentStr = d.SearchField.Text()
		}
	}
}

func (d *Display) saveDataToExcel() {
	columns := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}

	xlsx := excelize.NewFile()
	xlsx.SetSheetName("Sheet1", "Countries")
	xlsx.SetActiveSheet(1)

	for i := range columns {
		excelRow := 1
		for j := range data.Cached {
			// write only displayed rows/cards related countries
			if data.Cached[j].IsSearchedFor {
				res := ""
				switch columns[i] {
				case "A":
					res = data.Cached[j].Name.Common
				case "B":
					res = data.Cached[j].Name.Official
				case "C":
					res = data.Cached[j].Cca2
				case "D":
					res = data.Cached[j].Cca3
				case "E":
					res = data.Cached[j].Ccn3
				case "F":
					res = data.Cached[j].Cioc
				case "G":
					if data.Cached[j].Independent {
						res = "Yes"
					} else {
						res = "No"
					}
				case "H":
					res = data.Cached[j].Status
				case "I":
					if data.Cached[j].UNMember {
						res = "Yes"
					} else {
						res = "No"
					}
				case "J":
					if len(data.Cached[j].Capitals) > 0 {
						res = data.Cached[j].Capitals[0]
					} else {
						res = "N/A"
					}
				case "K":
					res = strconv.FormatFloat(data.Cached[j].Area, 'f', -1, 32)
				case "L":
					res = string(data.Cached[j].Population)
				}
				if err := xlsx.SetCellValue("Countries", columns[i]+strconv.Itoa(excelRow), res); err != nil {
					log.Fatalln(err)
				}
			}
			excelRow += 1
		}
	}
	if err := xlsx.SaveAs("output/general_info/excel/Countries.xlsx"); err != nil {
		log.Fatalln("error at excel save: ", err.Error())
	}
}

// init all continents with IsSelected set to false
func (d *Display) initContinents() {
	d.AllContinents = []string{"All", "Asia", "Africa", "Antarctica", "Europe", "Oceania", "North America", "South America"}
	for i := range d.AllContinents {
		d.Continents = append(d.Continents, Continent{Name: d.AllContinents[i]})
	}
}
