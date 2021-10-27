package globals

import "C"
import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"math/rand"
	"time"

	"image"
	"image/color"
	"log"
	"os"
	"strings"
)

func ColoredArea(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()
	clip.Rect{Max: size}.Add(gtx.Ops)
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

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

func GenerateRandomColor() color.NRGBA {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 255
	return color.NRGBA{
		R: uint8(rand.Intn(max-min) + min),
		G: uint8(rand.Intn(max-min) + min),
		B: uint8(rand.Intn(max-min) + min),
		A: 255,
	}
}

// isInJsonString - checks to see if the provided Input is a JSON string
//func isInJsonString(data []byte) bool {
//	var jsonMessage json.RawMessage
//	err := json.Unmarshal(data, &jsonMessage)
//	if err != nil {
//		return false
//	}
//	return true
//}
//
//// isInJson - checks to see if the provided input is a JSON (interface)
//func isInJson(data []byte) bool {
//	var jsonMessage map[string]interface{}
//	err := json.Unmarshal(data, &jsonMessage)
//	if err != nil {
//		return false
//	}
//	return true
//}
