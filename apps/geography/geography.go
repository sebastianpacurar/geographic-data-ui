package geography

import (
	"fmt"
	"gioui-experiment/apps"
	"gioui-experiment/apps/geography/data"
	"gioui-experiment/apps/geography/grid"
	"gioui-experiment/apps/geography/table"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"github.com/xuri/excelize/v2"
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
	}

	Display struct {
		// search
		SearchField component.TextField
		CurrentStr  string

		// layout buttons
		TableBtn   widget.Clickable
		GridBtn    widget.Clickable
		SaveAsXlsx widget.Clickable

		// grid and Table displays
		Grid  grid.Grid
		Table table.Table

		// grid or Table Selected display
		Selected interface{}
		Loaded   bool

		// api data
		Api data.Countries

		// slider
		Slider Slider

		// lock the context on the country
		ContextualSet bool

		// Contextual Viewed Country
		ContextualCountry data.Country
	}
)

func New(router *apps.Router) *Application {
	return &Application{
		Router: router,
	}
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (app *Application) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (app *Application) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Geography - countries, states, statistics",
	}
}

func (app *Application) LayoutView(gtx C, th *material.Theme) D {
	err := app.Display.Api.InitCountries()
	app.Display.FilterData()

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

	if app.Display.Selected == nil {
		app.Display.Selected = app.Display.Table
	}

	if err != nil {
		return material.H2(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	}

	app.Display.SearchField.SingleLine = true

	// run only once during lifetime
	if !app.Display.Loaded {
		for i := range data.Cached {
			data.Cached[i].Active = true
		}
		app.Display.Loaded = true
	}

	var dims D

	switch app.Display.Grid.Contextual.(type) {
	case data.Country:
		if !app.ContextualSet {
			app.Router.AppBar.SetContextualActions(
				[]component.AppBarAction{
					component.SimpleIconAction(&app.pinBtn, g.PinIcon,
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
			layout.Rigid(material.Body2(th, fmt.Sprintf("%f", app.Display.ContextualCountry.Area)).Layout),
		)
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
						app.Display.Slider.PushRight()
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
						app.Display.Slider.PushLeft()
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
						btn = material.IconButton(th, &app.Display.SaveAsXlsx, g.ExcelExportIcon, "Export to Excel")
						btn.Size = unit.Dp(20)
						btn.Inset = layout.UniformInset(unit.Dp(8))
						return btn.Layout(gtx)
					})
				})

				return layout.Flex{}.Layout(gtx,
					search,
					layout.Rigid(func(gtx C) D {
						return layout.Flex{Alignment: layout.End}.Layout(gtx,
							tblBtn,
							grdBtn,
							exportBtn,
						)
					}))
			}),

			// Selected display
			layout.Rigid(func(gtx C) D {
				return app.Display.Slider.Layout(gtx, layout.Horizontal, func(gtx C) D {
					var dims D
					switch app.Display.Selected.(type) {
					case grid.Grid:
						dims = app.Display.Grid.Layout(gtx, th)
					case table.Table:
						dims = app.Display.Table.Layout(gtx, th)
					}
					return dims
				})
			}))
	}
	return dims
}

func (app *Application) LayoutController(gtx C, th *material.Theme) D {
	return app.ControlPanel.Layout(gtx, th)
}

// FilterData - filter countries based on data.Cached and data.Cached manipulation
func (d *Display) FilterData() {
	if d.CurrentStr != d.SearchField.Text() {
		if d.SearchField.Len() > 0 {
			for i := range data.Cached {
				if strings.HasPrefix(strings.ToLower(data.Cached[i].Name.Common), strings.ToLower(d.SearchField.Text())) {
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
					res = fmt.Sprintf("%f", data.Cached[j].Area)
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
	if err := xlsx.SaveAs("./apps/geography/output/Countries.xlsx"); err != nil {
		log.Fatalln("error at excel save: ", err.Error())
	}
}
