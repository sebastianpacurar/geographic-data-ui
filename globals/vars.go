package globals

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"

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
		"dark-cyan":     {R: 0, G: 139, B: 139, A: 255},
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

func ColoredArea(gtx C, size image.Point, color color.NRGBA) D {
	defer op.Save(gtx.Ops).Load()
	clip.Rect{Max: size}.Add(gtx.Ops)
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return D{Size: size}
}
