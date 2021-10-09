package globals

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"

	"image"
	"image/color"
	"log"
	"os"
	"strings"
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

// GetAppsNames - What this does is to return all the folder names from "apps" package
// In case any of the results from reading the directory using Readdirnames(0) is a file,
// then do not append it to the apps []string.
// TODO: will be used with the menu and navigation drawer in the future
func GetAppsNames() []string {
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
	return apps
}
