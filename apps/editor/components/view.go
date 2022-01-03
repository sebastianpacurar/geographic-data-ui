package components

import (
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
)

type TextArea struct {
	field widget.Editor
	list  widget.List

	menu            component.MenuState
	ctxArea         component.ContextArea
	isMenuTriggered bool

	// menu options
	pasteBtn widget.Clickable
}

func (ta *TextArea) Layout(th *material.Theme) layout.FlexChild {
	ta.list.Axis = layout.Vertical
	ta.field.SingleLine = false
	ta.field.Alignment = text.Start

	if !ta.isMenuTriggered {
		var item component.MenuItemStyle
		item.LabelInset = layout.Inset{
			Top:    unit.Dp(5),
			Right:  unit.Dp(5),
			Bottom: unit.Dp(5),
			Left:   unit.Dp(5),
		}
		item = component.MenuItem(th, &ta.pasteBtn, "Paste")
		ta.menu = component.MenuState{
			Options: []func(gtx C) D{
				item.Layout,
			},
		}
	}

	return layout.Flexed(1, func(gtx C) D {
		border := widget.Border{
			Color:        g.Colours[colors.GREY],
			CornerRadius: unit.Dp(5),
			Width:        unit.Px(2),
		}
		switch {
		case ta.field.Focused():
			border.Color = th.Palette.ContrastBg
			border.Width = unit.Px(2)
		}

		return g.Inset.Layout(gtx, func(gtx C) D {
			return layout.Stack{}.Layout(gtx,
				layout.Stacked(func(gtx C) D {
					gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(gtx.Constraints.Max))
					return border.Layout(gtx, func(gtx C) D {
						return material.List(th, &ta.list).Layout(gtx, 1, func(gtx C, _ int) D {
							return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
								ed := material.Editor(th, &ta.field, "Type your Thoughts...")
								ed.SelectionColor = g.Colours[colors.TEXT_SELECTION]

								if ta.pasteBtn.Clicked() {
									ed.Editor.SetText(g.ClipBoardVal)
								}
								return ed.Layout(gtx)
							})
						})
					})
				}),
				layout.Expanded(func(gtx C) D {
					return ta.ctxArea.Layout(gtx, func(gtx C) D {
						gtx.Constraints.Min = image.Point{}
						return component.Menu(th, &ta.menu).Layout(gtx)
					})
				}))
		})
	})
}
