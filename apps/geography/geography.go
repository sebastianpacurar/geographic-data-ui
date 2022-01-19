package geography

import (
	"fmt"
	"gioui-experiment/apps"
	"gioui-experiment/apps/geography/data"
	"gioui-experiment/apps/geography/grid"
	"gioui-experiment/apps/geography/table"
	"gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
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
		th *material.Theme
		ControlPanel
		Display
		pinBtn widget.Clickable
		*apps.Router

		DisableCPBtn widget.Clickable
		btnStyle     material.ButtonStyle
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

		// FilterBy - used to search countries
		FilterBy string

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
		slider       Slider
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
		app.FilterData(table.NAME)

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
						data.Cached[j].ActiveContinent = true
					}
					break
				}
			}

			// set all countries to Active
			for i := range data.Cached {
				data.Cached[i].Active = true
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

			dims = layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Body2(th, app.Display.ContextualCountry.Cca2).Layout),
				layout.Rigid(material.Body2(th, app.Display.ContextualCountry.Cca3).Layout),
				layout.Rigid(material.Body2(th, app.Display.ContextualCountry.Ccn3).Layout),
				layout.Rigid(material.Body2(th, strconv.FormatFloat(app.Display.ContextualCountry.Area, 'f', -1, 32)).Layout),
				layout.Rigid(func(gtx C) D {
					return layout.Flex{}.Layout(gtx,
						layout.Flexed(1, func(gtx C) D {
							return widget.Image{
								Src: paint.NewImageOp(app.Display.ContextualCountry.FlagImg),
								Fit: widget.Contain,
							}.Layout(gtx)
						}))
				}))

		case nil:
			dims = layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {

					// search field
					search := layout.Flexed(1, func(gtx C) D {
						return app.Display.SearchField.Layout(gtx, th, "Search country")
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
								return app.ContinentsList.Layout(gtx, len(app.AllContinents), func(gtx C, i int) D {
									var (
										dim D
										btn material.ButtonStyle
									)
									btn = material.Button(th, &app.Continents[i].Btn, app.Continents[i].Name)
									btn.CornerRadius = unit.Dp(1)
									btn.Inset = layout.UniformInset(unit.Dp(10))
									btn.Background = globals.Colours[colours.WHITE]
									btn.Color = globals.Colours[colours.BLACK]
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

										// update ActiveContinent state
										for j := range data.Cached {
											if name == "All" {
												data.Cached[j].ActiveContinent = true
											} else {
												continents := strings.Join(data.Cached[j].Continents, " ")
												if strings.Contains(continents, app.Continents[i].Name) {
													data.Cached[j].ActiveContinent = true
												} else {
													data.Cached[j].ActiveContinent = false
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
													return globals.ColoredArea(gtx, size, globals.Colours[colours.AERO_BLUE])
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
																	return globals.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, 3), globals.Colours[colours.SEA_GREEN])
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
							dims = app.Display.Table.Layout(gtx, th, app.FilterBy)
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

// FilterData - filter countries based on data.Cached and data.Cached manipulation
func (d *Display) FilterData(FilterBy string) {
	if d.CurrentStr != d.SearchField.Text() {
		if d.SearchField.Len() > 0 {
			for i := range data.Cached {
				var res string

				switch FilterBy {
				case table.NAME:
					res = data.Cached[i].Name.Common
				case table.OFFICIAL_NAME:
					res = data.Cached[i].Name.Official
				case table.CAPITAL:
					res = data.Cached[i].Capital[0]
				}

				if strings.HasPrefix(strings.ToLower(res), strings.ToLower(d.SearchField.Text())) {
					data.Cached[i].Active = true
				} else {
					data.Cached[i].Active = false
				}
			}
			d.CurrentStr = d.SearchField.Text()
		} else if d.SearchField.Len() == 0 {
			for i := range data.Cached {
				data.Cached[i].Active = true
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
			if data.Cached[j].Active {
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
					if len(data.Cached[j].Capital) > 0 {
						res = data.Cached[j].Capital[0]
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
	if err := xlsx.SaveAs("output/geography/excel/Countries.xlsx"); err != nil {
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
