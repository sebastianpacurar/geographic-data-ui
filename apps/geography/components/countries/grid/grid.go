package grid

import (
	"gioui-experiment/apps/geography/components/countries/data"
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
)

type (
	Grid struct {
		cards []Card
		list  widget.List
		outlay.GridWrap
	}
)

func (gr *Grid) Layout(gtx C, th *material.Theme) D {
	gr.Alignment = layout.End
	gr.list.Axis = layout.Vertical
	gr.list.Alignment = layout.Middle

	for i := range data.Data {
		var capital string
		if len(data.Data[i].Capital) >= 1 {
			capital = data.Data[i].Capital[0]
		} else {
			capital = "N/A"
		}

		gr.cards = append(gr.cards, Card{
			Name:    data.Data[i].Name.Common,
			Capital: capital,
			Cioc:    data.Data[i].Cioc,
			FlagSrc: data.Data[i].FlagSrc.Png,
		})
	}

	return material.List(th, &gr.list).Layout(gtx, 1, func(gtx C, j int) D {
		return gr.GridWrap.Layout(gtx, len(data.Data), func(gtx C, i int) D {
			return globals.Inset.Layout(gtx, func(gtx C) D {
				return gr.cards[i].LayCard(gtx, th, &gr.cards[i])
			})
		})
	})
}
