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

type Geometry struct {
	ops     *op.Ops
	pos     f32.Point
	color   color.NRGBA
	squares int
	lastPos image.Point
}

// Square - Draw a Blue square based on size in cm and custom colo4
func (g *Geometry) Square(gtx C, size int, color color.NRGBA) D {
	g.squares++
	size *= globals.CM * size
	defer op.Save(gtx.Ops).Load()
	clip.Rect{Max: image.Pt(size*g.squares, size*g.squares)}.Add(gtx.Ops)
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return D{
		Size: image.Pt(size, size),
	}
}

// Layout - get dimensions of the square
func (g *Geometry) Layout(gtx C) D {
	return layout.Flex{}.Layout(
		gtx,
		layout.Rigid(func(gtx C) D {
			return g.Square(gtx, 2, globals.Colours["blue"])
		}),
	)
}
