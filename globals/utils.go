package globals

import "C"
import (
	"gioui.org/f32"
	"gioui.org/layout"
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
	dims := image.Rectangle{Max: gtx.Constraints.Max}
	defer clip.Rect(dims).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func RColoredArea(gtx layout.Context, size image.Point, r float32, color color.NRGBA) layout.Dimensions {
	bounds := f32.Rect(0, 0, float32(size.X), float32(size.Y))
	defer clip.RRect{
		Rect: bounds,
		SE:   r,
		SW:   r,
		NW:   r,
		NE:   r,
	}.Push(gtx.Ops).Pop()
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
