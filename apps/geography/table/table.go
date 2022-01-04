package table

import (
	"gioui-experiment/apps/geography/data"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Table struct {
		rows       []Row
		rowList    widget.List
		columnList widget.List
		loaded     bool
	}
)

func (t *Table) Layout(gtx C, th *material.Theme) D {
	t.rowList.Axis = layout.Vertical
	t.rowList.Alignment = layout.Middle
	t.columnList.Axis = layout.Horizontal
	t.columnList.Alignment = layout.Middle

	if !t.loaded {
		for i := range data.Data {
			t.rows = append(t.rows, Row{
				Name:         data.Data[i].Name.Common,
				OfficialName: data.Data[i].Name.Official,
				Capital:      data.Data[i].Capital,
				Independent:  data.Data[i].Independent,
				Status:       data.Data[i].Status,
				UNMember:     data.Data[i].UNMember,
				Cca2:         data.Data[i].Cca2,
				Cca3:         data.Data[i].Cca3,
				Ccn3:         data.Data[i].Ccn3,
				Area:         data.Data[i].Area,
				Population:   data.Data[i].Population,
				Active:       data.Data[i].Active,
			})
		}
		t.loaded = true
	} else {
		for i := range data.Data {
			t.rows[i].Active = data.Data[i].Active
			t.rows[i].Selected = data.Data[i].Selected
			t.rows[i].IsCPViewed = data.Data[i].IsCPViewed
		}
	}

	return material.List(th, &t.columnList).Layout(gtx, 1, func(gtx C, _ int) D {
		return material.List(th, &t.rowList).Layout(gtx, len(data.Data), func(gtx C, i int) D {
			if t.rows[i].Active {
				if t.rows[i].Click.Clicked() {
					if t.rows[i].Selected {
						data.Data[i].Selected = false
					} else {
						data.Data[i].Selected = true
					}
				}
			}
			return t.rows[i].LayRow(gtx, th)
		})
	})
}
