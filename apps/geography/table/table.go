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
		for i := range data.Cached {
			t.rows = append(t.rows, Row{
				Name:         data.Cached[i].Name.Common,
				OfficialName: data.Cached[i].Name.Official,
				Capital:      data.Cached[i].Capital,
				Independent:  data.Cached[i].Independent,
				Status:       data.Cached[i].Status,
				UNMember:     data.Cached[i].UNMember,
				Cca2:         data.Cached[i].Cca2,
				Cca3:         data.Cached[i].Cca3,
				Ccn3:         data.Cached[i].Ccn3,
				Area:         data.Cached[i].Area,
				Population:   data.Cached[i].Population,
				Active:       data.Cached[i].Active,
			})
		}
		t.loaded = true
	} else {
		for i := range data.Cached {
			t.rows[i].Active = data.Cached[i].Active
			t.rows[i].Selected = data.Cached[i].Selected
			t.rows[i].IsCPViewed = data.Cached[i].IsCPViewed
		}
	}

	return material.List(th, &t.columnList).Layout(gtx, 1, func(gtx C, _ int) D {
		return material.List(th, &t.rowList).Layout(gtx, len(data.Cached), func(gtx C, i int) D {
			var dims D
			if t.rows[i].Active {
				if t.rows[i].Click.Clicked() {
					if t.rows[i].Selected {
						data.Cached[i].Selected = false
					} else {
						data.Cached[i].Selected = true
					}
				}
				dims = t.rows[i].LayRow(gtx, th)
			}
			return dims
		})
	})
}
