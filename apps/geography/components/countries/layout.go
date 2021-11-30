package countries

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type (
	Display struct {
		table
		data.Countries
	}

	table struct {
		rowList, columnList widget.List
		rows                []tableRow
	}

	tableRow struct {
		name  string
		list  widget.List
		cells []tableCell
	}

	tableCell struct {
		name   string
		border widget.Border
		layout func(gtx C) D
	}
)

func (d *Display) Layout(gtx C, th *material.Theme) D {
	err := d.InitCountries()
	if err != nil {
		return material.H2(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	}

	d.table.rowList.Axis = layout.Vertical
	return material.List(th, &d.table.rowList).Layout(gtx, len(data.Data), func(gtx C, rowIndex int) D {
		return material.Body1(th, data.Data[rowIndex].Name.Common).Layout(gtx)
	})
}
