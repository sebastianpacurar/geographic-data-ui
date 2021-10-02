package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
	"log"
	"os"
	"strconv"
)

var defaultMargin = unit.Dp(10)

type C = layout.Context
type D = layout.Dimensions

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

type counter struct {
	count    int
	increase widget.Clickable
	decrease widget.Clickable
	reset    widget.Clickable
}

type UI struct {
	theme   *material.Theme
	counter counter
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

func (c *counter) Layout(th *material.Theme, gtx C) D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		layout.Flexed(1, func(gtx C) D {
			return layout.Center.Layout(
				gtx,
				material.H2(th, strconv.Itoa(c.count)).Layout,
			)
		}),

		layout.Rigid(
			layout.Spacer{
				Height: defaultMargin,
			}.Layout,
		),

		layout.Flexed(2, func(gtx C) D {
			for range c.increase.Clicks() {
				c.count++
			}
			btn := material.Button(th, &c.increase, "Increase")
			btn.Background = color.NRGBA{G: 255, A: 255}
			return btn.Layout(gtx)
		}),

		layout.Rigid(
			layout.Spacer{
				Height: defaultMargin,
			}.Layout,
		),

		layout.Flexed(2, func(gtx C) D {
			for range c.decrease.Clicks() {
				c.count--
			}
			btn := material.Button(th, &c.decrease, "Decrease")
			btn.Background = color.NRGBA{R: 255, A: 255}
			return btn.Layout(gtx)
		}),

		layout.Rigid(
			layout.Spacer{
				Height: defaultMargin,
			}.Layout,
		),

		layout.Flexed(2, func(gtx C) D {
			btn := material.Button(th, &c.reset, "Reset")
			for range c.reset.Clicks() {
				c.count = 0
			}
			return btn.Layout(gtx)
		}),
	)
}

func (ui *UI) Layout(gtx C) D {
	inset := layout.UniformInset(defaultMargin)
	return inset.Layout(
		gtx,
		func(gtx C) D {
			return ui.counter.Layout(ui.theme, gtx)
		})
}
