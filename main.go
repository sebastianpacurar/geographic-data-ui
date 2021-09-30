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
	"log"
	"os"
)

type C = layout.Context
type D = layout.Dimensions

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Title of the app"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		if err := draw(w); err != nil {
			log.Fatal(err) // this will return os.Exit(1), killing the program
		}

		os.Exit(0) // keep the app running

	}()
	app.Main()
}

func draw(w *app.Window) error {
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
					func(gtx C) D {
						// define margins around the button
						margins := layout.Inset{
							Right: unit.Dp(35),
							Left:  unit.Dp(35),
						}

						// return the margins (lay them on the screen)
						return margins.Layout(
							gtx,
							// return the button which is within the margins.
							func(gtx C) D {
								plusBtn := material.Button(th, &plusButton, "Plus")
								return plusBtn.Layout(gtx)
							})
					},
				),
				layout.Rigid(
					func(gtx C) D {
						margins := layout.Inset{
							Right: unit.Dp(35),
							Left:  unit.Dp(35),
						}
						return margins.Layout(
							gtx,
							func(gtx C) D {
								minusBtn := material.Button(th, &minusButton, "Minus")
								return minusBtn.Layout(gtx)
							})
					}),
			)
			e.Frame(gtx.Ops)

		case system.DestroyEvent:
			return e.Err // this will be sent before the app receives close signal
		}
	}
	return nil
}
