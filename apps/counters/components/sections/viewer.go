package sections

import (
	"gioui-experiment/apps/counters/components/controllers"
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
	"strconv"
)

type View struct{}

func (v *View) Layout(th *material.Theme, gtx controllers.C) controllers.D {
	size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
	return layout.Stack{
		Alignment: layout.NW,
	}.Layout(
		gtx,
		layout.Expanded(func(gtx controllers.C) controllers.D {
			view := globals.ColoredArea(
				gtx,
				gtx.Constraints.Constrain(size),
				globals.Colours["antique-white"],
			)
			return view
		}),

		layout.Stacked(func(gtx controllers.C) controllers.D {
			mainVal := material.H3(th, strconv.FormatInt(globals.Count, 10))
			if globals.Count < 0 {
				mainVal.Color = globals.Colours["red"]
			} else if globals.Count > 0 {
				mainVal.Color = globals.Colours["dark-green"]
			} else {
				mainVal.Color = globals.Colours["grey"]
			}

			return layout.Inset{
				Top:    unit.Dp(20),
				Right:  unit.Dp(50),
				Bottom: unit.Dp(20),
				Left:   unit.Dp(50),
			}.Layout(gtx, func(gtx controllers.C) controllers.D {
				return mainVal.Layout(gtx)
			})
		}))
}
