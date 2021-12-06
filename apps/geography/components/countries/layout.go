package countries

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
	"image"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type (
	Display struct {
		cards []card
		data.Countries
		grid outlay.GridWrap
		list widget.List
	}

	card struct {
		name, capital string
		cioc, flagSrc string
		click         widget.Clickable
		flag          image.Image
	}
)

func (d *Display) LayCard(gtx C, th *material.Theme, card *card) D {
	size := image.Pt(150, 200)
	gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return widget.Border{
				Color:        g.Colours[colors.GREY],
				CornerRadius: unit.Dp(18),
				Width:        unit.Px(2),
			}.Layout(gtx, func(gtx C) D {
				return g.RColoredArea(gtx,
					size,
					unit.Dp(18),
					g.Colours[colors.WHITE],
				)
			})

		}),
		layout.Stacked(func(gtx C) D {
			return g.Inset.Layout(gtx, func(gtx C) D {
				gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(size))
				return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceAround}.Layout(gtx,

					// country name
					layout.Rigid(func(gtx C) D {
						return layout.Flex{}.Layout(gtx,
							layout.Flexed(1, func(gtx C) D {
								return layout.Center.Layout(gtx, func(gtx C) D {
									return material.Body2(th, card.name).Layout(gtx)
								})
							}),
						)
					}),

					// TODO: fix this or find a workaround
					// (capital area) country flag (temporary broken)
					layout.Rigid(func(gtx C) D {
						//country.flag = d.processFlagFromURL(country)
						return layout.Flex{}.Layout(gtx,
							layout.Flexed(1, func(gtx C) D {
								return layout.Center.Layout(gtx, func(gtx C) D {
									return material.Body2(th, card.capital).Layout(gtx)
								})
							}))
					}))
			})
		}))

}

func (d *Display) Layout(gtx C, th *material.Theme) D {
	err := d.InitCountries()
	if err != nil {
		return material.H2(th, fmt.Sprintf("Error when fetching countries: %s", err)).Layout(gtx)
	}
	d.grid.Axis = layout.Horizontal
	d.grid.Alignment = layout.End
	d.list.Axis = layout.Vertical
	d.list.Alignment = layout.Middle

	for i := range data.Data {
		var capital string
		if len(data.Data[i].Capital) >= 1 {
			capital = data.Data[i].Capital[0]
		} else {
			capital = "N/A"
		}

		d.cards = append(d.cards, card{
			name:    data.Data[i].Name.Common,
			capital: capital,
			cioc:    data.Data[i].Cioc,
			flagSrc: data.Data[i].FlagSrc.Png,
		})
	}

	return material.List(th, &d.list).Layout(gtx, 1, func(gtx C, j int) D {
		return d.grid.Layout(gtx, len(data.Data), func(gtx C, i int) D {
			return g.Inset.Layout(gtx, func(gtx C) D {
				return material.Clickable(gtx, &d.cards[i].click, func(gtx C) D {
					for _, e := range gtx.Events(d) {
						switch e := e.(type) {
						case pointer.Event:
						}
					}
					return d.LayCard(gtx, th, &d.cards[i])
				})
			})
		})
	})
}

//func (d *Display) processFlagFromURL(card *card) image.Image {
//	r, re := http.Get(card.flagSrc)
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
//			log.Fatalln(fmt.Sprintf("error when closing response body reader for %s", card.name))
//		}
//	}(r.Body)
//
//	res, _ := png.Decode(r.Body)
//	return res
//}
