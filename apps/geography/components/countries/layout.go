package countries

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	"gioui-experiment/apps/geography/components/countries/grid"
	"gioui-experiment/apps/geography/components/countries/table"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/outlay"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type (
	Display struct {
		// search
		searchField component.TextField

		// layout buttons
		tableBtn widget.Clickable
		gridBtn  widget.Clickable

		// api data
		data.Countries

		// grid and table
		grid.Card
		cards []grid.Card
		grid  outlay.GridWrap
		list  widget.List
		table table.Table
	}
)

func (d *Display) Layout(gtx C, th *material.Theme) D {
	err := d.InitCountries()
	if err != nil {
		return material.H2(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	}
	d.grid.Alignment = layout.End
	d.list.Axis = layout.Vertical
	d.list.Alignment = layout.Middle

	d.searchField.SingleLine = true

	// TODO: isolate in a gofile named "grid.go" or smthing
	for i := range data.Data {
		var capital string
		if len(data.Data[i].Capital) >= 1 {
			capital = data.Data[i].Capital[0]
		} else {
			capital = "N/A"
		}

		d.cards = append(d.cards, grid.Card{
			Name:    data.Data[i].Name.Common,
			Capital: capital,
			Cioc:    data.Data[i].Cioc,
			FlagSrc: data.Data[i].FlagSrc.Png,
		})
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, func(gtx C) D {
					return d.searchField.Layout(gtx, th, "Search country")
				}),

				layout.Rigid(func(gtx C) D {
					return layout.Flex{Alignment: layout.End}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(10),
								Right:  unit.Dp(8),
								Bottom: unit.Dp(8),
								Left:   unit.Dp(8),
							}.Layout(gtx, func(gtx C) D {
								return material.Button(th, &d.tableBtn, "Table").Layout(gtx)
							})
						}),
						layout.Rigid(func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(10),
								Right:  unit.Dp(10),
								Bottom: unit.Dp(8),
								Left:   unit.Dp(8),
							}.Layout(gtx, func(gtx C) D {
								return material.Button(th, &d.gridBtn, "Grid").Layout(gtx)
							})
						}))
				}),
			)
		}),

		// TABLE
		layout.Rigid(func(gtx C) D {
			return d.table.Layout(gtx, th)
		}))

	// GRID
	// TODO: isolate in the same "grid.go" file
	//return material.List(th, &d.list).Layout(gtx, 1, func(gtx C, j int) D {
	//	return d.grid.Layout(gtx, len(data.Data), func(gtx C, i int) D {
	//		return g.Inset.Layout(gtx, func(gtx C) D {
	//			return d.LayCard(gtx, th, &d.cards[i])
	//		})
	//	})
	//})
}

// TODO: currently stuck
//func (d *Display) processFlagFromURL(grid *grid) image.Image {
//	r, re := http.Get(grid.flagSrc)
//	if re != nil {
//		path := "apps/geography/components/countries/assets/placeholders/no_flag.png"
//		f, fe := os.Open(path)
//		if fe != nil {
//			log.Fatalln("error when opening no_flag.png")
//		}
//		defer func(f *os.File) {
//			err := f.Close()
//			if err != nil {
//				log.Fatalln("error closing the no_flag.png reader")
//			}
//		}(f)
//
//		res, _ := png.Decode(f)
//		return res
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			log.Fatalln(fmt.Sprintf("error when closing response body reader for %s", grid.name))
//		}
//	}(r.Body)
//
//	res, _ := png.Decode(r.Body)
//	return res
//}
