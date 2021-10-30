package globals

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"image/color"
)

var (
	MenuWidth         = unit.Dp(225)
	CountersMenuWidth = unit.Dp(250)
	DefaultMargin     = unit.Dp(10)
	Colours           = map[string]color.NRGBA{
		"red":             {R: 255, A: 255},
		"dark-red":        {R: 139, A: 255},
		"blue":            {B: 255, A: 255},
		"green":           {G: 255, A: 255},
		"dark-green":      {G: 180, A: 255},
		"dark-cyan":       {G: 139, B: 139, A: 255},
		"sea-green":       {R: 46, G: 139, B: 87, A: 255},
		"deep-sky-blue":   {R: 0, G: 191, B: 255, A: 255},
		"dark-slate-grey": {R: 47, G: 79, B: 79, A: 255},
		"grey":            {R: 128, G: 128, B: 128, A: 255},
		"white":           {R: 255, G: 255, B: 255, A: 255},
		"black":           {A: 255},
		"antique-white":   {R: 250, G: 235, B: 215, A: 255},
		"aero-blue":       {R: 201, G: 255, B: 229, A: 255},
	}
	Inset   = layout.UniformInset(DefaultMargin)
	SpacerX = layout.Rigid(
		layout.Spacer{
			Width: DefaultMargin,
		}.Layout,
	)
)
