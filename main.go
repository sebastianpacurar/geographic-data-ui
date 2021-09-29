package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Title of the app"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)
		var ops op.Ops                               // operations from the UI
		var plusButton widget.Clickable              // plusButton is a clickable widget
		var minusButton widget.Clickable             // minusButton is another clickable widget
		th := material.NewTheme(gofont.Collection()) // th defines the material design style

		// listen for events in the window.
		for e := range w.Events() {
			// detect what type of event
			switch e := e.(type) {
			// this is sent when the application should re-render.
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceEvenly,
				}.Layout(
					gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							plusBtn := material.Button(th, &plusButton, "Plus")
							return plusBtn.Layout(gtx)
						},
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							minusBtn := material.Button(th, &minusButton, "Minus")
							return minusBtn.Layout(gtx)
						},
					),
					//layout.Rigid(
					//	layout.Spacer{
					//		Height: unit.Dp(25),
					//	}.Layout,
					//),
				)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
