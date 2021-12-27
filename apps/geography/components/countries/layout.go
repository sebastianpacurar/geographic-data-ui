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

		// grid and table displays
		grid  grid.Grid
		table table.Table

		// grid or table selected display
		selected interface{}

		// api data
		data.Countries

		// slider
		slider Slider
	}
)

func (d *Display) Layout(gtx C, th *material.Theme) D {
	err := d.InitCountries()
	if d.selected == nil {
		d.selected = d.grid
	}

	if err != nil {
		return material.H2(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	}

	d.searchField.SingleLine = true

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, func(gtx C) D {
					return d.searchField.Layout(gtx, th, "Search country")
				}),

				layout.Rigid(func(gtx C) D {
					return layout.Flex{Alignment: layout.End}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
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
						}),

						layout.Rigid(func(gtx C) D {
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
