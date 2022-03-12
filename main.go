package main

import (
	application "gioui-experiment/apps"
	"gioui-experiment/apps/general_info"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"log"
	"os"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Geography Application"),
			app.Size(unit.Dp(1400), unit.Dp(900)),
		)
		if err := Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func Run(w *app.Window) error {
	var ops op.Ops
	th := material.NewTheme(gofont.Collection())

	router := application.NewRouter()
	router.Register(1, general_info.New(&router))

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				router.Layout(gtx, th)
				e.Frame(gtx.Ops)
			case system.DestroyEvent:
				return e.Err
			}
		}
	}
}
