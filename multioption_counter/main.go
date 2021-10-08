package main

import (
	comp "gioui-experiment/multioption_counter/components"
	"gioui-experiment/multioption_counter/globals"

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

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	ui := NewUI()
	go func() {
		w := app.NewWindow(
			app.Title("Multi Option Counter"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)
		if err := ui.Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type UI struct {
	theme      *material.Theme
	counter    comp.Counter
	startValue comp.StartValue
}

func (ui *UI) Run(w *app.Window) error {
	var ops op.Ops
	for event := range w.Events() {
		switch event := event.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, event)
			ui.Layout(gtx)
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

func NewUI() *UI {
	ui := &UI{}
	ui.theme = material.NewTheme(gofont.Collection())
	return ui
}

func (ui *UI) Layout(gtx C) D {
	inset := layout.UniformInset(globals.DefaultMargin)
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		layout.Flexed(0.112, func(gtx C) D {
			return inset.Layout(gtx, func(gtx C) D {
				return ui.startValue.Layout(ui.theme, gtx)
			})
		}),
		layout.Flexed(1, func(gtx C) D {
			return inset.Layout(gtx, func(gtx C) D {
				return ui.counter.Layout(ui.theme, gtx)
			})
		}),
	)
}
