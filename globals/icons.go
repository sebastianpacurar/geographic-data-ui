package globals

import (
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

// PlusIcon - used for the plusBtn
var PlusIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ContentAdd)
	return icon
}()

// MinusIcon - used for the minusBtn
var MinusIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ContentRemove)
	return icon
}()

var MenuIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationMenu)
	return icon
}()
