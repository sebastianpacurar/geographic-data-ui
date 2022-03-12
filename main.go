package main

import (
	application "gioui-experiment/sections"
	"gioui-experiment/sections/covid_stats"
	"gioui-experiment/sections/general_info"
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

type UI struct {
	th     *material.Theme
	ops    op.Ops
	router application.Router
}

func NewUI() *UI {
	r := application.NewRouter()
	r.Register(1, general_info.New(&r))
	r.Register(2, covid_stats.New(&r))

	return &UI{
		th:     material.NewTheme(gofont.Collection()),
		router: r,
	}
}

func main() {
	geoApp := NewUI()

	go func() {
		w := app.NewWindow(
			app.Title("Geographic Data"),
			app.Size(unit.Dp(1400), unit.Dp(900)),
		)
		if err := geoApp.Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func (ui *UI) Run(w *app.Window) error {
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ui.ops, e)
				ui.router.Layout(gtx, ui.th)
				e.Frame(gtx.Ops)
			case system.DestroyEvent:
				return e.Err
			}
		}
	}
}
