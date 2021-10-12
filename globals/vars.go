package globals

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"image/color"
	"math"
)

var (
	CM            = int(math.Floor(37.795275591))
	Count         = int64(0)
	CountUnit     = int64(1)
	ResetVal      = int64(0)
	DefaultMargin = unit.Dp(10)
	DefaultBorder = widget.Border{
		Color:        Colours["grey"],
		CornerRadius: unit.Dp(3),
		Width:        unit.Px(2),
	}
	Colours = map[string]color.NRGBA{
		"red":             {R: 255, A: 255},
		"dark-red":        {R: 139, A: 255},
		"blue":            {B: 255, A: 255},
		"green":           {G: 255, A: 255},
		"dark-green":      {G: 180, A: 255},
		"dark-cyan":       {G: 139, B: 139, A: 255},
		"dark-slate-grey": {R: 47, G: 79, B: 79, A: 255},
		"grey":            {R: 128, G: 128, B: 128, A: 255},
		"white":           {R: 255, G: 255, B: 255, A: 255},
		"black":           {A: 255},
		"antique-white":   {R: 250, G: 235, B: 215, A: 255},
	}
	Inset   = layout.UniformInset(DefaultMargin)
	SpacerX = layout.Rigid(
		layout.Spacer{
			Width: DefaultMargin,
		}.Layout,
	)
	SpacerY = layout.Rigid(
		layout.Spacer{
			Height: DefaultMargin,
		}.Layout,
	)
)
