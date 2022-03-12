package globals

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

/*
/// Colors
*/

const (
	FLAME_RED       = "flame-red"
	SEA_GREEN       = "sea-green"
	LIGHT_SEA_GREEN = "light-sea-green"
	GREY            = "grey"
	WHITE           = "white"
	BLACK           = "black"
	ANTIQUE_WHITE   = "antique-white"
	AERO_BLUE       = "aero-blue"
	LIGHT_SALMON    = "light-salmon"
	NYANZA          = "nyanza"
	LAVENDERBLUSH   = "lavenderblush"
	ELECTRIC_BLUE   = "electric-blue"
	LIGHT_YELLOW    = "light-yellow"

	CP_RESIZER     = "cp_resizer"
	TEXT_SELECTION = "text_selection"
	CARD_COLOR     = "card_color"
)

/*
/// Global Vars
*/

var (
	ClipBoardVal string
	Colours      = map[string]color.NRGBA{
		"red":       {R: 255, A: 255},
		"flame-red": {R: 220, G: 85, B: 44, A: 255},
		"blue":      {B: 255, A: 255},
		"green":     {G: 255, A: 255},

		"white":           {R: 255, G: 255, B: 255, A: 255},
		"lavenderblush":   {R: 255, G: 240, B: 245, A: 255},
		"antique-white":   {R: 250, G: 235, B: 215, A: 255},
		"light-yellow":    {R: 255, G: 255, B: 152, A: 255},
		"aero-blue":       {R: 201, G: 255, B: 229, A: 255},
		"light-salmon":    {R: 255, G: 207, B: 188, A: 255},
		"nyanza":          {R: 233, G: 255, B: 219, A: 255},
		"sea-green":       {R: 46, G: 139, B: 87, A: 255},
		"light-sea-green": {R: 32, G: 178, B: 170, A: 255},
		"deep-sky-blue":   {R: 0, G: 191, B: 255, A: 255},
		"electric-blue":   {R: 125, G: 249, B: 255, A: 255},

		"black":           {A: 255},
		"dark-slate-grey": {R: 40, G: 70, B: 70, A: 175},
		"grey":            {R: 128, G: 128, B: 128, A: 255},
		"dark-red":        {R: 139, A: 255},
		"dark-green":      {G: 180, A: 255},
		"dark-cyan":       {G: 139, B: 139, A: 255},

		"text_selection": {R: 191, G: 255, B: 209, A: 255},
		"card_color":     {R: 237, G: 237, B: 237, A: 255},
		"cp_resizer":     {R: 40, G: 70, B: 70, A: 175},
	}
)

/*
/// Geometry related
*/

func ColoredArea(gtx C, size image.Point, color color.NRGBA) D {
	dims := image.Rectangle{Max: gtx.Constraints.Max}
	paint.FillShape(gtx.Ops, color, clip.Rect(dims).Op())
	return D{Size: size}
}

func RColoredArea(gtx C, size image.Point, r float32, color color.NRGBA) D {
	bounds := f32.Rect(0, 0, float32(size.X), float32(size.Y))
	paint.FillShape(gtx.Ops, color, clip.UniformRRect(bounds, r).Op(gtx.Ops))
	return D{Size: size}
}

/*
/// Icons
*/

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
