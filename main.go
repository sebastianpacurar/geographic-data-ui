package main

import (
	application "gioui-experiment/apps"
	"gioui-experiment/apps/counters"
	"gioui-experiment/apps/editor"
	"gioui-experiment/apps/geography"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
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
	router.Register(0, counters.New(&router))
	router.Register(1, editor.New(&router))
	router.Register(2, geography.New(&router))

	for event := range w.Events() {
		switch event := event.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, event)
			router.Layout(gtx, th)
			event.Frame(gtx.Ops)
		case key.Event:
			switch event.Name {
			case key.NameEscape:
				return nil
			}
		case system.DestroyEvent:
			return event.Err
		}
	}
	return nil
}
