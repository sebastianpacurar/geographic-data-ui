package countries

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	"gioui-experiment/apps/geography/components/countries/grid"
	"gioui-experiment/apps/geography/components/countries/table"
	g "gioui-experiment/globals"
	"gioui.org/layout"
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
)

type (
	Display struct {
		// search
		searchField component.TextField
		currentStr  string

		// layout buttons
		tableBtn   widget.Clickable
		gridBtn    widget.Clickable
		saveAsXlsx widget.Clickable

		// grid and table displays
		grid  grid.Grid
		table table.Table

		// grid or table selected display
		selected interface{}
		loaded   bool

		// api data
		api data.Countries

		// slider
		slider Slider
	}
)

func (d *Display) Layout(gtx C, th *material.Theme) D {
	err := d.api.InitCountries()
	d.filterData()
	if d.selected == nil {
		d.selected = d.table
	}

	if err != nil {
		return material.H2(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	}

	d.searchField.SingleLine = true

	// run only once during lifetime
	if !d.loaded {
		for i := range data.Data {
			data.Data[i].Active = true
		}
		d.loaded = true
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {

			// search field
			search := layout.Flexed(1, func(gtx C) D {
				return d.searchField.Layout(gtx, th, "Search country")
			})

			// table button
			tblBtn := layout.Rigid(func(gtx C) D {
				if d.tableBtn.Clicked() {
					d.selected = d.table
					d.slider.PushRight()
				}
				switch d.selected.(type) {
				case table.Table:
					gtx = gtx.Disabled()
				}
				return layout.Inset{
					Top:    unit.Dp(10),
					Right:  unit.Dp(8),
					Bottom: unit.Dp(8),
					Left:   unit.Dp(8),
				}.Layout(gtx, func(gtx C) D {
					return material.Button(th, &d.tableBtn, "Table").Layout(gtx)
				})
			})

			// grid button
			grdBtn := layout.Rigid(func(gtx C) D {
				if d.gridBtn.Clicked() {
					d.selected = d.grid
					d.slider.PushLeft()
				}
				switch d.selected.(type) {
				case grid.Grid:
					gtx = gtx.Disabled()
				}
				return layout.Inset{
					Top:    unit.Dp(10),
					Right:  unit.Dp(8),
					Bottom: unit.Dp(8),
					Left:   unit.Dp(8),
				}.Layout(gtx, func(gtx C) D {
					return material.Button(th, &d.gridBtn, "Grid").Layout(gtx)
				})
			})

			//Export to excel button
			exportBtn := layout.Rigid(func(gtx C) D {
				if d.saveAsXlsx.Clicked() {
					columns := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}

					xlsx := excelize.NewFile()
					xlsx.SetSheetName("Sheet1", "Countries")
					xlsx.SetActiveSheet(1)

					for i := range columns {
						for j := range data.Data {
							// write only displayed rows/cards related countries
							if data.Data[i].Active {
								res := ""
								switch columns[i] {
								case "A":
									res = data.Data[j].Name.Common
								case "B":
									res = data.Data[j].Name.Official
								case "C":
									res = data.Data[j].Cca2
								case "D":
									res = data.Data[j].Cca3
								case "E":
									res = data.Data[j].Ccn3
								case "F":
									res = data.Data[j].Cioc
								case "G":
									if data.Data[j].Independent {
										res = "Yes"
									} else {
										res = "No"
									}
								case "H":
									res = data.Data[j].Status
								case "I":
									if data.Data[j].UNMember {
										res = "Yes"
									} else {
										res = "No"
									}
								case "J":
									if len(data.Data[j].Capital) > 0 {
										res = data.Data[j].Capital[0]
									} else {
										res = "N/A"
									}
								case "K":
									res = fmt.Sprintf("%f", data.Data[j].Area)
								case "L":
									res = string(data.Data[j].Population)
								}
								if err := xlsx.SetCellValue("Countries", columns[i]+strconv.Itoa(j+1), res); err != nil {
									log.Fatalln(err)
								}
							}
						}
					}
					if err := xlsx.SaveAs("./apps/geography/output/Countries.xlsx"); err != nil {
						log.Fatalln("error at excel save: ", err.Error())
					}
				}
				return layout.Inset{
					Top:    unit.Dp(10),
					Right:  unit.Dp(10),
					Bottom: unit.Dp(8),
					Left:   unit.Dp(8),
				}.Layout(gtx, func(gtx C) D {
					var btn material.IconButtonStyle
					btn = material.IconButton(th, &d.saveAsXlsx, g.ExcelExportIcon, "Export to Excel")
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
			return d.slider.Layout(gtx, func(gtx C) D {
				switch d.selected.(type) {
				case grid.Grid:
					return d.grid.Layout(gtx, th)
				case table.Table:
					return d.table.Layout(gtx, th)
				}
				return D{}
			})
		}))
}

// filterData - filter countries based on data.Data and data.Cached manipulation
func (d *Display) filterData() {
	if d.currentStr != d.searchField.Text() {
		if d.searchField.Len() > 0 {
			for i := range data.Data {
				if strings.HasPrefix(strings.ToLower(data.Data[i].Name.Common), strings.ToLower(d.searchField.Text())) {
					data.Data[i].Active = true
				} else {
					data.Data[i].Active = false
				}
			}
			d.currentStr = d.searchField.Text()
		} else if d.searchField.Len() == 0 {
			for i := range data.Data {
				data.Data[i].Active = true
			}
			d.currentStr = d.searchField.Text()
		}
	}
}
