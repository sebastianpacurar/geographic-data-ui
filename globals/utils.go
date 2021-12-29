package globals

import "C"
import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
	"image/color"
)

func ColoredArea(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	dims := image.Rectangle{Max: gtx.Constraints.Max}
	defer clip.Rect(dims).Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, color)
	return layout.Dimensions{Size: size}
}

func RColoredArea(gtx layout.Context, size image.Point, r unit.Value, color color.NRGBA) layout.Dimensions {
	bounds := f32.Rect(0, 0, float32(size.X), float32(size.Y))
	defer clip.RRect{
		Rect: bounds,
		SE:   float32(gtx.Px(r)),
		SW:   float32(gtx.Px(r)),
		NW:   float32(gtx.Px(r)),
		NE:   float32(gtx.Px(r)),
	}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, color)
	return layout.Dimensions{Size: size}
}
