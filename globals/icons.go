package globals

import (
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var PlusIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ContentAdd)
	return icon
}()

var RefreshIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationRefresh)
	return icon
}()

var MinusIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ContentRemove)
	return icon
}()

var MenuIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationMenu)
	return icon
}()

var ExcelExportIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.FileFileDownload)
	return icon
}()

var PinIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.MapsPinDrop)
	return icon
}()

var VisibilityIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.ActionVisibility)
	return icon
}()
