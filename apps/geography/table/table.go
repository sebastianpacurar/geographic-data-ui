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
		headerList layout.List
		rowList    widget.List
		columnList widget.List
		loaded     bool
	}
)

func (t *Table) Layout(gtx C, th *material.Theme, searchBy string) D {
	t.rowList.Axis = layout.Vertical
	t.rowList.Alignment = layout.Middle
	t.columnList.Axis = layout.Horizontal
	t.columnList.Alignment = layout.Middle

	if !t.loaded {
		for i := range data.Cached {
			t.rows = append(t.rows, Row{
				Name:            data.Cached[i].Name.Common,
				OfficialName:    data.Cached[i].Name.Official,
				Capital:         data.Cached[i].Capital,
				Region:          data.Cached[i].Region,
				Subregion:       data.Cached[i].Subregion,
				Languages:       data.Cached[i].Languages,
				IddRoot:         data.Cached[i].Idd.Root,
				IddSuffixes:     data.Cached[i].Idd.Suffixes,
				TopLevelDomains: data.Cached[i].TopLevelDomain,
				Independent:     data.Cached[i].Independent,
				Status:          data.Cached[i].Status,
				UNMember:        data.Cached[i].UNMember,
				Landlocked:      data.Cached[i].Landlocked,
				Cca2:            data.Cached[i].Cca2,
				Cca3:            data.Cached[i].Cca3,
				Ccn3:            data.Cached[i].Ccn3,
				Cioc:            data.Cached[i].Cioc,
				Fifa:            data.Cached[i].Fifa,
				Area:            data.Cached[i].Area,
				Population:      data.Cached[i].Population,
				Latitude:        data.Cached[i].LatLng[0],
				Longitude:       data.Cached[i].LatLng[1],
				Continents:      data.Cached[i].Continents,
				StartOfWeek:     data.Cached[i].StartOfWeek,
				CarSigns:        data.Cached[i].Car.Signs,
				CarSide:         data.Cached[i].Car.Side,

				Active:          data.Cached[i].Active,
				ActiveContinent: data.Cached[i].ActiveContinent,
			})
		}
		t.loaded = true
	} else {
		for i := range data.Cached {
			t.rows[i].Active = data.Cached[i].Active
			t.rows[i].ActiveContinent = data.Cached[i].ActiveContinent
			t.rows[i].Selected = data.Cached[i].Selected
			t.rows[i].IsCPViewed = data.Cached[i].IsCPViewed
		}
	}

	return material.List(th, &t.columnList).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,

			// Header Area
			layout.Rigid(func(gtx C) D {
				return t.headerList.Layout(gtx, 1, func(gtx C, i int) D {
					return t.rows[i].LayRow(gtx, th, true)
				})
			}),

			// Row Area
			layout.Flexed(1, func(gtx C) D {
				return material.List(th, &t.rowList).Layout(gtx, len(data.Cached), func(gtx C, i int) D {
					var dims D
					if t.rows[i].Active && t.rows[i].ActiveContinent {
						if t.rows[i].Click.Clicked() {
							if t.rows[i].Selected {
								data.Cached[i].Selected = false
							} else {
								data.Cached[i].Selected = true
							}
						}
						dims = t.rows[i].LayRow(gtx, th, false)
					}
					return dims
				})
			}))
	})

}
