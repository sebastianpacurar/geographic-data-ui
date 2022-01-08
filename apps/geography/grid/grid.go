package grid

import (
	"encoding/json"
	"gioui-experiment/apps/geography/data"
	g "gioui-experiment/globals"
	"gioui.org/io/clipboard"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Grid struct {
		cards      []Card
		list       widget.List
		wrap       outlay.GridWrap
		Contextual interface{}
		loaded     bool
	}
)

func (gr *Grid) Layout(gtx C, th *material.Theme) D {

	gr.wrap.Alignment = layout.End
	gr.list.Axis = layout.Vertical
	gr.list.Alignment = layout.Middle

	if !gr.loaded {
		for i := range data.Cached {
			gr.cards = append(gr.cards, Card{
				Name:     data.Cached[i].Name.Common,
				Active:   data.Cached[i].Active,
				Selected: data.Cached[i].Selected,
				Flag:     data.Cached[i].FlagImg,
			})
		}
		gr.loaded = true
	} else {
		for i := range data.Cached {
			gr.cards[i].Active = data.Cached[i].Active
			gr.cards[i].Selected = data.Cached[i].Selected
		}
	}

	return material.List(th, &gr.list).Layout(gtx, 1, func(gtx C, _ int) D {
		return gr.wrap.Layout(gtx, len(data.Cached), func(gtx C, i int) D {
			var content D

			// copy only this specific card
			if gr.cards[i].copyToClipBtn.Clicked() {
				res, _ := json.MarshalIndent(data.Cached[i], "", "\t")
				clipboard.WriteOp{
					Text: string(res),
				}.Add(gtx.Ops)
				g.ClipBoardVal = string(res)
			}

			if gr.cards[i].viewBtn.Clicked() {
				gr.Contextual = data.Cached[i] // interface to assert type when enabling ContextualAppBar
				data.Cached[i].IsCtxtActive = true
				op.InvalidateOp{}.Add(gtx.Ops)
			}

			if gr.cards[i].selectBtn.Clicked() {
				data.Cached[i].Selected = true
			} else if gr.cards[i].deselectBtn.Clicked() {
				data.Cached[i].Selected = false
			}

			if gr.cards[i].Active {
				content = layout.UniformInset(unit.Dp(15)).Layout(gtx, func(gtx C) D {
					return gr.cards[i].LayCard(gtx, th)
				})
			}
			return content
		})
	})
}
