package globals

import (
	"gioui.org/layout"
	"gioui.org/unit"

	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

var (
	Count         = int64(0)
	ResetVal      = int64(0)
	DefaultMargin = unit.Dp(10)
	Colours       = map[string]color.NRGBA{
		"red":           {R: 255, A: 255},
		"blue":          {B: 255, A: 255},
		"green":         {G: 255, A: 255},
		"dark-green":    {G: 180, A: 255},
		"grey":          {R: 128, G: 128, B: 128, A: 255},
		"white":         {R: 255, G: 255, B: 255, A: 255},
		"black":         {A: 255},
		"antique-white": {R: 250, G: 235, B: 215, A: 255},
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
