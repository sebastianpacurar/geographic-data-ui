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

type TopBar struct {
	menuBtn widget.Clickable
	title   string
}

// Layout - is composed of a Stack layout which returns the first dimension as
// the fullWidth of the X-Axis, and 120 pixels on Y-Axis
func (ab *TopBar) Layout(gtx C) D {
	return layout.Stack{}.Layout(gtx,
		// Expand the colored area, allowing for child Stacked widgets to overlap its dimensions
		layout.Expanded(func(gtx C) D {
			// size is a bit hackish - X = full width; Y = the smallest Stacked sibling + defaultMargin/2
			// otherwise, moving the app window on a larger monitor, causes the Y size to grow,
			// and the other way around.
			size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y+10/2)
			bar := globals.ColoredArea(
				gtx,
				gtx.Constraints.Constrain(size),
				globals.Colours["dark-cyan"],
			)
			return bar
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
