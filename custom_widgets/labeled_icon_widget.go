package custom_widgets

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

type LabeledIconBtn struct {
	Theme      *material.Theme
	LabelColor color.NRGBA
	Button     *widget.Clickable
	Icon       *widget.Icon
	Label      string
}

func (lib LabeledIconBtn) Layout(gtx layout.Context) layout.Dimensions {
	// Set the TextSize for the Label
	lib.Theme.TextSize.Scale(14.0 / 16.0)

	// return a ButtonLayout dimension which contains the Icon and Label
	return material.ButtonLayout(lib.Theme, lib.Button).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(globals.DefaultMargin).Layout(gtx, func(gtx layout.Context) layout.Dimensions {

			labeledIcon := layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
			}
			spacer := unit.Dp(5)
			// This is the actual Icon of the Button. It will return the dimensions of the Icon widget
			layIcon := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Right: spacer,
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					var d layout.Dimensions
					if lib.Icon != nil {
						size := gtx.Px(unit.Dp(56)) - 2*gtx.Px(unit.Dp(16))
						gtx.Constraints = layout.Exact(image.Pt(size, size))
						d = lib.Icon.Layout(gtx, lib.LabelColor)
					}
					return d
				})
			})

			// This is the actual Label of the Button, treated as a Body1 Typography element
			layLabel := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: spacer}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					l := material.Body1(lib.Theme, lib.Label)
					l.Color = lib.LabelColor
					return l.Layout(gtx)
				})
			})

			// eventually return labeledIcon with all its children added as parameters in the right order
			return labeledIcon.Layout(gtx, layIcon, layLabel)
		})
	})
}
