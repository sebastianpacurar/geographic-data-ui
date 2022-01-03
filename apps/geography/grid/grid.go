package grid

import (
	"encoding/json"
	"gioui-experiment/apps/geography/data"
	g "gioui-experiment/globals"
	"gioui.org/io/clipboard"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Grid struct {
		cards  []Card
		list   widget.List
		wrap   outlay.GridWrap
		loaded bool
	}
)

func (gr *Grid) Layout(gtx C, th *material.Theme) D {

	gr.wrap.Alignment = layout.End
	gr.list.Axis = layout.Vertical
	gr.list.Alignment = layout.Middle

	if !gr.loaded {
		for i := range data.Data {
			gr.cards = append(gr.cards, Card{
				Name:     data.Data[i].Name.Common,
				Cca2:     data.Data[i].Cca2,
				Active:   data.Data[i].Active,
				Selected: data.Data[i].Selected,
			})
		}
		gr.loaded = true
	} else {
		for i := range data.Data {
			gr.cards[i].Active = data.Data[i].Active
			gr.cards[i].Selected = data.Data[i].Selected
		}
	}

	return material.List(th, &gr.list).Layout(gtx, 1, func(gtx C, _ int) D {
		return gr.wrap.Layout(gtx, len(data.Data), func(gtx C, i int) D {
			var content D

			// copy only this specific card
			if gr.cards[i].copyToClipBtn.Clicked() {
				res, _ := json.MarshalIndent(data.Data[i], "", "\t")
				clipboard.WriteOp{
					Text: string(res),
				}.Add(gtx.Ops)
				g.ClipBoardVal = string(res)
			}

			if gr.cards[i].Click.Clicked() {
				//apps.Router{}
			}

			if gr.cards[i].selectBtn.Clicked() {
				data.Data[i].Selected = true
			} else if gr.cards[i].deselectBtn.Clicked() {
				data.Data[i].Selected = false
			}

			if gr.cards[i].Active {
				content = layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx C) D {
					return gr.cards[i].LayCard(gtx, th)
				})
			}
			return content
		})
	})
}
