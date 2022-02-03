package playground

import (
	"gioui-experiment/apps/playground/controllers"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	ControlPanel struct {
		controllers []Controller
		list        widget.List

		test      controllers.TestController
		testState component.DiscloserState
	}

	Controller struct {
		name   string
		layout func(C, *Controller) D
	}
)

func (cp *ControlPanel) Layout(gtx C, th *material.Theme) D {
	cp.list.Axis = layout.Vertical
	cp.controllers = []Controller{
		{
			name: "Nothing for now",
			layout: func(gtx C, c *Controller) D {
				return component.SimpleDiscloser(th, &cp.testState).Layout(gtx,
					material.Body1(th, c.name).Layout,
					func(gtx C) D {
						return controllerInset.Layout(gtx, func(gtx C) D {
							return cp.test.Layout(gtx, th)
						})
					})
			},
		},
	}
	// return a vertical list of (discloser, divider) groups, as ListElements
	return material.List(th, &cp.list).Layout(gtx, len(cp.controllers), func(gtx C, i int) D {
		return cp.controllers[i].layout(gtx, &cp.controllers[i])
	})
}

// LayOutset - wraps the discloser and divider in a vertical flex layout
func (cp *ControlPanel) LayOutset(gtx C, discloser, divider layout.FlexChild) D {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, discloser, divider)
}
