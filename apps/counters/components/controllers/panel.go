package controllers

import (
	"gioui-experiment/apps/counters/components/controllers/control_panel"
	"gioui-experiment/custom_themes/colors"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
)

type (
	ControlPanel struct {
		vh       control_panel.ValueHandler
		inc      control_panel.Incrementor
		incState component.DiscloserState
		vhState  component.DiscloserState
	}
)

var controllerInset = layout.Inset{
	Top:    unit.Dp(10),
	Right:  unit.Dp(20),
	Bottom: unit.Dp(10),
	Left:   unit.Dp(5),
}

func (cp *ControlPanel) Layout(gtx C, th *material.Theme) D {
	incControl := layout.Rigid(func(gtx C) D {
		return component.SimpleDiscloser(th, &cp.incState).Layout(gtx,
			material.Body1(th, "Manual Incrementors").Layout,
			func(gtx C) D {
				return controllerInset.Layout(gtx, func(gtx C) D {
					return cp.inc.Layout(gtx, th)
				})
			})
	})

	vhControl := layout.Rigid(func(gtx C) D {
		return component.SimpleDiscloser(th, &cp.vhState).Layout(gtx,
			material.Body1(th, "Start and Step Values").Layout,
			func(gtx C) D {
				return controllerInset.Layout(gtx, func(gtx C) D {
					return cp.vh.Layout(gtx, th)
				})
			})
	})

	return layout.Stack{Alignment: layout.NW}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return g.ColoredArea(gtx, gtx.Constraints.Max, g.Colours[colors.AERO_BLUE])
		}),
		layout.Stacked(func(gtx C) D {
			containerSize := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
			gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(containerSize))
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, incControl, vhControl)
		}))
}
