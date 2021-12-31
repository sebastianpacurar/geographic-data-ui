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

//var LockCLosedIcon = func() *widget.Icon {
//	icon, _ := widget.NewIcon(icons.ActionLock)
//	return icon
//}()
//
//var LockOpenedIcon = func() *widget.Icon {
//	icon, _ := widget.NewIcon(icons.ActionLockOpen)
//	return icon
//}()

var ExcelExportIcon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.FileFileDownload)
	return icon
}()
