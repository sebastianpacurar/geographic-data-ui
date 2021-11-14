package sections

import (
	"gioui-experiment/apps/counters/components/controllers"
	"gioui-experiment/apps/counters/components/controllers/control_panel"
	"gioui-experiment/apps/counters/components/data"
	"gioui-experiment/custom_themes/colors"
	g "gioui-experiment/globals"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
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
	cp := layout.Rigid(func(gtx C) D {
		return v.ControlPanel.Layout(gtx, th)
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

				/// 100 height BAR DISPLAYED ABOVE THE DISPLAYED NUMBER
				///
				layout.Stacked(func(gtx C) D {
					return layout.Stack{}.Layout(gtx,
						layout.Expanded(func(gtx C) D {
							defer clip.UniformRRect(f32.Rectangle{
								Max: layout.FPt(image.Pt(gtx.Constraints.Max.X, 100)),
							}, 0).Push(gtx.Ops).Pop()
							return D{}
						}))
				}),

				/// DISPLAYED NUMBER
				///
				layout.Stacked(func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(80),
						Right:  unit.Dp(50),
						Bottom: unit.Dp(20),
						Left:   unit.Dp(50),
					}.Layout(gtx, func(gtx C) D {
						var val string
						seq := cv.GetActiveSequence()
						switch seq {
						case data.PRIMES, data.FIBS:
							val = strconv.FormatUint(cv.Cache[seq][cv.Index], 10)
						case data.NATURALS, data.INTEGERS:
							val = strconv.FormatUint(cv.Displayed, 10)
						}
						return material.H5(th, val).Layout(gtx)
					})
				}))
		}),

		///
		/// CONTROL PANEL IS RENDERED HERE
		cp,
	)
}
