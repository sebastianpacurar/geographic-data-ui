package general_info

import (
	"fmt"
	"gioui-experiment/apps"
	"gioui-experiment/apps/general_info/data"
	"gioui-experiment/apps/general_info/grid"
	"gioui-experiment/apps/general_info/table"
	"gioui-experiment/apps/general_info/views"
	"gioui-experiment/globals"
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

	Application struct {
		*apps.Router
		th *material.Theme

		ControlPanel
		Display
		cv views.CountryView

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

func New(router *apps.Router) *Application {
	return &Application{
		Router: router,
		th:     material.NewTheme(gofont.Collection()),
	}
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Tag: &app.DisableCPBtn,
			},
			Layout: func(gtx C, bg, fg color.NRGBA) D {
				var lbl string
				if app.DisableCPBtn.Clicked() {
					app.isCPDisabled = !app.isCPDisabled
				}
				if !app.isCPDisabled {
					lbl = "Disable CP"
				} else {
					lbl = "Enable CP"
				}
				return material.Button(app.th, &app.DisableCPBtn, lbl).Layout(gtx)
			},
		},
	}
}

func (app *Application) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (app *Application) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Geography - countries, states, statistics",
	}
}

func (app *Application) IsCPDisabled() bool {
	return app.isCPDisabled
}

func (app *Application) LayoutView(gtx C, th *material.Theme) D {
	err := app.Api.InitCountries()
	if err != nil {
		// in case of no internet connection, or if the main request fails
		return material.H6(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	} else {
		app.SearchByColumn(table.SearchBy)

		// run only once at start
		if !app.initialSetup {
			app.SearchField.SingleLine = true

			// initialize all flags
			data.ProcessFlags()

			// initialize Table View at start
			app.Selected = app.Display.Table

			// initialize continents to "All" as Selected
			app.initContinents()
			for i := range app.Continents {
				if app.Continents[i].Name == "All" {
					app.Continents[i].IsSelected = true
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
			app.initialSetup = true
		}

		for _, e := range app.AppBar.Events(gtx) {
			switch e.(type) {
			case component.AppBarContextMenuDismissed:
				app.Display.ContextualSet = false
				app.Display.Grid.Contextual = nil
				for i := range data.Cached {
					if data.Cached[i].IsCtxtActive {
						data.Cached[i].IsCtxtActive = false
					}
				}
				op.InvalidateOp{}.Add(gtx.Ops)

			case component.AppBarNavigationClicked:
				app.ModalNavDrawer.Appear(gtx.Now)
				app.NavAnim.Disappear(gtx.Now)
			}
		}

		var dims D

		switch app.Grid.Contextual.(type) {
		case data.Country:
			if !app.ContextualSet {
				app.Router.AppBar.SetContextualActions(
					[]component.AppBarAction{
						component.SimpleIconAction(&app.pinBtn, globals.PinIcon,
							component.OverflowAction{
								Name: "Pin Country",
								Tag:  &app.pinBtn,
							},
						),
					},
					[]component.OverflowAction{},
				)

				for i := range data.Cached {
					if data.Cached[i].IsCtxtActive {
						app.Display.ContextualCountry = data.Cached[i]
					}
				}
				app.Router.AppBar.ToggleContextual(gtx.Now, app.Display.ContextualCountry.Name.Common)
				app.ContextualSet = true
				op.InvalidateOp{}.Add(gtx.Ops)
			}

			// layout Country View
			dims = app.cv.Layout(gtx, th, app.ContextualCountry)

		case nil:
			dims = layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {

					// search field
					search := layout.Flexed(1, func(gtx C) D {
						return app.Display.SearchField.Layout(gtx, th, fmt.Sprintf("Search by %s", table.SearchBy))
					})

					// Table button
					tblBtn := layout.Rigid(func(gtx C) D {
						if app.Display.TableBtn.Clicked() {
							app.Display.Selected = app.Display.Table
							app.Display.slider.PushRight()
						}
						switch app.Display.Selected.(type) {
						case table.Table:
							gtx = gtx.Disabled()
						}
						return layout.Inset{
							Top:    unit.Dp(10),
							Right:  unit.Dp(8),
							Bottom: unit.Dp(8),
							Left:   unit.Dp(8),
						}.Layout(gtx, material.Button(th, &app.Display.TableBtn, "Table").Layout)
					})

					// grid button
					grdBtn := layout.Rigid(func(gtx C) D {
						if app.Display.GridBtn.Clicked() {
							app.Display.Selected = app.Display.Grid
							app.Display.slider.PushLeft()
						}
						switch app.Display.Selected.(type) {
						case grid.Grid:
							gtx = gtx.Disabled()
						}
						return layout.Inset{
							Top:    unit.Dp(10),
							Right:  unit.Dp(8),
							Bottom: unit.Dp(8),
							Left:   unit.Dp(8),
						}.Layout(gtx, material.Button(th, &app.Display.GridBtn, "Grid").Layout)
					})

					//Export to excel button
					exportBtn := layout.Rigid(func(gtx C) D {
						if app.Display.SaveAsXlsx.Clicked() {

						}
						return layout.Inset{
							Top:    unit.Dp(10),
							Right:  unit.Dp(10),
							Bottom: unit.Dp(8),
							Left:   unit.Dp(8),
						}.Layout(gtx, func(gtx C) D {
							var btn material.IconButtonStyle
							btn = material.IconButton(th, &app.Display.SaveAsXlsx, globals.ExcelExportIcon, "Export to Excel")
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
								return app.ContinentsList.Layout(gtx, len(app.AllContinents), func(gtx C, i int) D {
									var (
										dim D
										btn material.ButtonStyle
									)
									btn = material.Button(th, &app.Continents[i].Btn, app.Continents[i].Name)
									btn.CornerRadius = unit.Dp(1)
									btn.Inset = layout.UniformInset(unit.Dp(10))
									btn.Background = globals.Colours[globals.WHITE]
									btn.Color = globals.Colours[globals.BLACK]
									dim = btn.Layout(gtx)

									if app.Continents[i].Btn.Clicked() {
										name := app.Continents[i].Name
										app.Continents[i].IsSelected = true

										// update the tab ui
										for j := range app.Continents {
											if name != app.Continents[j].Name {
												app.Continents[j].IsSelected = false
											}
										}

										// update IsActiveContinent state
										for j := range data.Cached {
											if name == "All" {
												data.Cached[j].IsActiveContinent = true
											} else {
												continents := strings.Join(data.Cached[j].Continents, " ")
												if strings.Contains(continents, app.Continents[i].Name) {
													data.Cached[j].IsActiveContinent = true
												} else {
													data.Cached[j].IsActiveContinent = false
												}
											}
										}
										op.InvalidateOp{}.Add(gtx.Ops)
									}

									if app.Continents[i].IsSelected {
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
															lbl = material.Body1(th, app.Continents[i].Name)
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
					return app.Display.slider.Layout(gtx, layout.Horizontal, func(gtx C) D {
						var dims D
						switch app.Display.Selected.(type) {
						case grid.Grid:
							dims = app.Display.Grid.Layout(gtx, th)
						case table.Table:
							dims = app.Display.Table.Layout(gtx, th)
						}
						return dims
					})
				}),
			)
		}
		return dims
	}
}

func (app *Application) LayoutController(gtx C, th *material.Theme) D {
	return app.ControlPanel.Layout(gtx, th)
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
