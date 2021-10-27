package countries

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Display struct{}

func (d *Display) Layout(th *material.Theme, gtx C) D {
	data := make([]string, len(Data))
	for i := range data {
		data[i] = Data[i].Name.Common
	}
	fmt.Println(data)
	return material.H2(th, data[0]).Layout(gtx)
}
