package app_layout

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/x/component"
	"image"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type AppBar struct {
	menuBtn widget.Clickable
	title   string
	Apps    []string
}

// Layout - is composed of a Stack layout which returns the first dimension as
// the fullWidth of the X-Axis, and 120 pixels on Y-Axis
func (ab *AppBar) Layout(gtx C) D {
	if len(ab.Apps) == 0 {
		ab.Apps = globals.GetAppsNames()
	}
	fullWidth := gtx.Constraints.Max.X
	return layout.Stack{}.Layout(gtx,
		// Expand the colored area, allowing for child Stacked widgets to overlap its dimensions
		layout.Expanded(func(gtx C) D {

			// TODO: Fix issue with image.PT for Y-Axis. If moving the app on a bigger monitor, the AppBar grows
			return globals.ColoredArea(gtx, image.Pt(fullWidth, 120), globals.Colours["dark-cyan"])
		}),
		// This returns a Stacked layout which returns a custom Inset, which will eventually
		// return the Menu Button as a SimpleIconButton
		layout.Stacked(func(gtx C) D {
			// using a hackish Inset, so the Menu Button can be perfectly aligned with the other widgets
			return layout.Inset{
				Left: unit.Dp(10),
				Top:  unit.Dp(5),
			}.Layout(gtx, func(gtx C) D {
				btn := component.SimpleIconButton(
					globals.Colours["dark-cyan"],
					globals.Colours["white"],
					&ab.menuBtn,
					globals.MenuIcon,
				)
				return btn.Layout(gtx)
			})
		}),
	)
}