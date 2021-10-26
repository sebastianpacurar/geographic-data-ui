package globals

import (
	"gioui-experiment/apps/counters/components/utils"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"image/color"
)

var (
	CounterVals = &utils.CurrentValues{
		Enabled:    true,
		CurrVal:    "signed",
		Count:      0,
		UCount:     0,
		CountUnit:  1,
		UCountUnit: 1,
		ResetVal:   0,
		UResetVal:  0,
		Primes: utils.Primes{
			PEnabled:   false,
			PCurrIndex: 0,
		},
		Fibs: utils.Fibs{
			FEnabled:   false,
			FCurrIndex: 0,
		},
	}

	MenuWidth     = unit.Dp(225)
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
		"sea-green":       {R: 46, G: 139, B: 87, A: 255},
		"deep-sky-blue":   {R: 0, G: 191, B: 255, A: 255},
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
)
