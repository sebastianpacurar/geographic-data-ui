package grid

import (
	"gioui.org/widget"
	"image"
)

type (
	Card struct {
		Name, Capital string
		Cioc, FlagSrc string
		clicked       bool
		Click         widget.Clickable
		flag          image.Image
	}
)
