package globals

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"image"
	"image/color"
	"log"
	"os"
	"strings"
)

var (
	Count         = int64(0)
	ResetVal      = int64(0)
	DefaultMargin = unit.Dp(10)
	Colours       = map[string]color.NRGBA{
		"red":             {R: 255, A: 255},
		"dark-red":        {R: 139, A: 255},
		"blue":            {B: 255, A: 255},
		"green":           {G: 255, A: 255},
		"dark-green":      {G: 180, A: 255},
		"dark-cyan":       {G: 139, B: 139, A: 255},
		"dark-slate-grey": {R: 47, G: 79, B: 79, A: 255},
		"grey":            {R: 128, G: 128, B: 128, A: 255},
		"white":           {R: 255, G: 255, B: 255, A: 255},
		"black":           {A: 255},
		"antique-white":   {R: 250, G: 235, B: 215, A: 255},
	}
	Inset   = layout.UniformInset(DefaultMargin)
	SpacerX = layout.Rigid(
		layout.Spacer{
			Width: DefaultMargin,
		}.Layout,
	)
	SpacerY = layout.Rigid(
		layout.Spacer{
			Height: DefaultMargin,
		}.Layout,
	)
)

// ColoredArea - returns the dimensions of a customizable coloured area
// args - graphics context, a size for x and y of type int, and a color
func ColoredArea(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()
	clip.Rect{Max: size}.Add(gtx.Ops)
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

// GetAppNames - What this does is to return all the folder names from "apps" package
// In case any of the results from reading the directory using Readdirnames(0) is a file,
// then do not append it to the apps []string.
// TODO: will be used with the menu and navigation drawer in the future
func GetAppNames() ([]string, error) {
	var apps []string
	file, err := os.Open("apps")
	if err != nil {
		log.Fatalf("issue with the location of \"apps\" directory: %s", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatalf("issue when closing the file reader: %s", err)
		}
	}(file)
	appNames, err := file.Readdirnames(0)
	if err != nil {
		log.Fatalf("issue when reading from \"apps\" directory: %s", err)
	}
	for _, s := range appNames {
		if !strings.Contains(s, ".go") {
			apps = append(apps, s)
		}
	}
	return apps, err
}
