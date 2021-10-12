package geometry

import (
	"gioui-experiment/globals"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Draw func()

type Geometry struct {
	Color  color.NRGBA
	coords []float32
	Draw   func(C, []float32, color.NRGBA) D
	drawer Draw
}

func (G Geometry) FRect(gtx C, geom []float32, color color.NRGBA) D {
	defer op.Save(gtx.Ops).Load()
	coords := f32.Rect(geom[0], geom[1], geom[2], geom[3])
	clip.RRect{Rect: coords}.Add(gtx.Ops)
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return D{
		Size: image.Pt(int(coords.Size().X), int(coords.Size().Y)),
	}
}

func (G Geometry) Rect(gtx C, size image.Point, color color.NRGBA) D {
	defer op.Save(gtx.Ops).Load()
	clip.Rect{Max: size}.Add(gtx.Ops)
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return D{
		Size: size,
	}
}

func (G *Geometry) Layout(gtx C) D {
	coords := []float32{float32(gtx.Constraints.Max.X), float32(gtx.Constraints.Max.Y), 0, 0}
	return layout.Stack{}.Layout(
		gtx,
		layout.Expanded(func(gtx C) D {
			screen := G.FRect(gtx, coords, globals.Colours["antique-white"])
			return screen
		}),
		layout.Stacked(func(gtx C) D {
			return G.FRect(gtx, []float32{500, 500, 100, 100}, globals.Colours["black"])
		}),
	)
}
