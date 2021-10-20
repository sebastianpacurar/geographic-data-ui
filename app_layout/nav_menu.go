package app_layout

import (
	"gioui.org/widget"
)

type Menu struct {
	Active int
	Items  []MenuItem
}

type MenuItem struct {
	Name  string
	Click widget.Clickable
	W     func(item *MenuItem, gtx C) D
	Num   int
}
