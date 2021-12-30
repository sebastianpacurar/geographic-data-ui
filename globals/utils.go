package globals

import "C"
import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image"
	"image/color"
)

func ColoredArea(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	dims := image.Rectangle{Max: gtx.Constraints.Max}
	paint.FillShape(gtx.Ops, color, clip.Rect(dims).Op())
	return layout.Dimensions{Size: size}
}

func RColoredArea(gtx layout.Context, size image.Point, r float32, color color.NRGBA) layout.Dimensions {
	bounds := f32.Rect(0, 0, float32(size.X), float32(size.Y))
	paint.FillShape(gtx.Ops, color, clip.UniformRRect(bounds, r).Op(gtx.Ops))
	return layout.Dimensions{Size: size}
}
