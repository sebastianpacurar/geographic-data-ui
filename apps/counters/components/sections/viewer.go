package sections

import (
	"gioui-experiment/apps/counters/components/controllers"
	"gioui-experiment/apps/counters/components/controllers/control_panel"
	"gioui-experiment/apps/counters/components/data"
	"gioui-experiment/custom_themes/colors"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
	"strconv"
)

type View struct {
	inc          control_panel.Incrementor
	ControlPanel controllers.ControlPanel
}

func (v *View) Layout(gtx C, th *material.Theme) D {
	cv := data.CurrVals
	seq := cv.GetActiveSequence()

	/// CONTROL PANEL
	controlPanel := layout.Rigid(func(gtx C) D {
		return v.ControlPanel.Layout(gtx, th)
	})

	/// DISPLAYED NUMBER
	displayedNumber := layout.Rigid(func(gtx C) D {
		return layout.Inset{
			Top:    unit.Dp(10),
			Right:  unit.Dp(50),
			Bottom: unit.Dp(20),
			Left:   unit.Dp(50),
		}.Layout(gtx, func(gtx C) D {
			var val string
			switch seq {
			case data.PRIMES, data.FIBS:
				val = strconv.FormatUint(cv.Cache[seq][cv.Index], 10)
			case data.NATURALS, data.INTEGERS:
				val = strconv.FormatUint(cv.Displayed, 10)
			}
			return material.H5(th, val).Layout(gtx)
		})
	})

	size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			width := gtx.Constraints.Max.X - gtx.Px(g.CountersMenuWidth)
			containerSize := image.Pt(width, gtx.Constraints.Max.Y)
			gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))
			return layout.Stack{
				Alignment: layout.NW,
			}.Layout(gtx,
				layout.Expanded(func(gtx C) D {
					view := g.ColoredArea(
						gtx,
						gtx.Constraints.Constrain(size),
						g.Colours[colors.ANTIQUE_WHITE],
					)
					return view
				}),

				layout.Stacked(func(gtx C) D {
					containerSize := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
					gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						displayedNumber,
					)
				}),
			)
		}),
		controlPanel,
	)
}
