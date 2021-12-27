package grid

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"image"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Hovered struct {
		isHovered bool
	}
)

func (h *Hovered) Hovering(gtx C) bool {
	start := h.isHovered
	for _, ev := range gtx.Events(h) {
		switch ev := ev.(type) {
		case pointer.Event:
			switch ev.Type {
			case pointer.Enter:
				h.isHovered = true
			case pointer.Leave:
				h.isHovered = false
			case pointer.Cancel:
				h.isHovered = false
			}
		}
	}
	if h.isHovered != start {
		op.InvalidateOp{}.Add(gtx.Ops)
	}
	return h.isHovered
}

func (h *Hovered) Layout(gtx C) D {
	defer clip.Rect(image.Rectangle{Max: gtx.Constraints.Max}).Push(gtx.Ops).Pop()
	pointer.InputOp{
		Tag:   h,
		Types: pointer.Enter | pointer.Leave,
	}.Add(gtx.Ops)
	return D{Size: gtx.Constraints.Max}
}
