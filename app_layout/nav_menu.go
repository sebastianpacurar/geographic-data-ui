package app_layout

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
)

type Menu struct {
	Apps []string
}

func (m Menu) Layout(th *material.Theme, gtx C) D {
	//if len(m.Apps) == 0 {
	//	m.Apps = globals.GetAppsNames()
	//}
	return layout.Stack{
		Alignment: layout.NW,
	}.Layout(
		gtx,
		layout.Expanded(func(gtx C) D {
			size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
			bar := globals.ColoredArea(
				gtx,
				gtx.Constraints.Constrain(size),
				globals.Colours["sea-green"],
			)
			return bar
		}),

		layout.Stacked(func(gtx C) D {
			var appsList = &widget.List{
				List: layout.List{
					Axis: layout.Vertical,
				},
			}

			widgets := []layout.Widget{
				material.H4(th, "Menu").Layout,
				material.H6(th, "Counters").Layout,
				material.H6(th, "Geography").Layout,
			}

			return material.List(th, appsList).Layout(gtx, len(widgets), func(gtx C, i int) D {
				return layout.UniformInset(globals.DefaultMargin).Layout(gtx, widgets[i])
			})
		}),
	)

}
