package grid

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Grid struct {
		cards []Card
		list  widget.List
		outlay.GridWrap
		loaded bool
	}
)

func (gr *Grid) Layout(gtx C, th *material.Theme) D {

	gr.Alignment = layout.End
	gr.list.Axis = layout.Vertical
	gr.list.Alignment = layout.Middle

	if !gr.loaded {
		for i := range data.Data {
			gr.cards = append(gr.cards, Card{
				Name:   data.Data[i].Name.Common,
				Cca2:   data.Data[i].Cca2,
				Active: data.Data[i].Active,
			})
		}
		gr.loaded = true
	} else {
		for i := range data.Data {
			gr.cards[i].Active = data.Data[i].Active
		}
	}

	return material.List(th, &gr.list).Layout(gtx, 1, func(gtx C, _ int) D {
		return gr.GridWrap.Layout(gtx, len(data.Data), func(gtx C, i int) D {
			var content D

			if gr.cards[i].Click.Clicked() {
				fmt.Println(fmt.Sprintf("%s", gr.cards[i].Name))
			}

			if gr.cards[i].Click.Hovered() {
				data.Data[i].Hovered = true
			} else {
				data.Data[i].Hovered = false
			}

			if gr.cards[i].Active {
				content = g.Inset.Layout(gtx, func(gtx C) D {
					return gr.cards[i].LayCard(gtx, th, &gr.cards[i])
				})
			}
			return content
		})
	})
}
