package main

import (
	application "gioui-experiment/apps"
	"gioui-experiment/apps/editor"
	"gioui-experiment/apps/geography"
	"gioui-experiment/apps/playground"
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
			app.Title("Gio UI Experiment"),
			app.Size(unit.Dp(1000), unit.Dp(800)),
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
	router.Register("pg", playground.New(&router))
	router.Register(1, geography.New(&router))
	router.Register(2, editor.New(&router))

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
