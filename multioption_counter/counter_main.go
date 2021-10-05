package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	count         = 0
	defaultMargin = unit.Dp(10)
	colours       = map[string]color.NRGBA{
		"red":           {R: 255, A: 255},
		"blue":          {B: 255, A: 255},
		"green":         {G: 255, A: 255},
		"grey":          {R: 128, G: 128, B: 128, A: 255},
		"black":         {A: 255},
		"antique-white": {R: 250, G: 235, B: 215, A: 255},
	}
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

type valueInput struct {
	widget.Editor
	Invalid bool
	oldVal  string
}

func (vi valueInput) Updated() bool {
	newVal := vi.Editor.Text()
	updated := newVal != vi.oldVal
	vi.oldVal = newVal
	return updated
}

type valueChanger struct {
	inpVal    valueInput
	changeVal widget.Clickable
}

type counter struct {
	increase, decrease, reset widget.Clickable
}

type UI struct {
	theme        *material.Theme
	counter      counter
	valueChanger valueChanger
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
			currVal := material.H2(th, strconv.Itoa(count))
			if count < 0 {
				currVal.Color = colours["red"]
			} else if count > 0 {
				currVal.Color = colours["green"]
			} else {
				currVal.Color = colours["grey"]
			}
			return layout.Center.Layout(
				gtx,
				currVal.Layout,
			)
		}),

		layout.Rigid(
			layout.Spacer{
				Height: defaultMargin,
			}.Layout,
		),

		layout.Flexed(0.1, func(gtx C) D {
			for range c.increase.Clicks() {
				count++
			}
			btn := material.Button(th, &c.increase, "Increase")
			btn.Background = colours["green"]
			btn.Color = colours["black"]
			return btn.Layout(gtx)
		}),

		layout.Rigid(
			layout.Spacer{
				Height: defaultMargin,
			}.Layout,
		),

		layout.Flexed(0.1, func(gtx C) D {
			for range c.decrease.Clicks() {
				count--
			}
			btn := material.Button(th, &c.decrease, "Decrease")
			btn.Background = colours["red"]
			return btn.Layout(gtx)
		}),

		layout.Rigid(
			layout.Spacer{
				Height: defaultMargin,
			}.Layout,
		),

		layout.Flexed(0.1, func(gtx C) D {
			for range c.reset.Clicks() {
				count = 0
			}
			btn := material.Button(th, &c.reset, "Reset")
			btn.Background = colours["blue"]
			return btn.Layout(gtx)
		}),
	)
}

func (vc *valueChanger) Layout(th *material.Theme, gtx C) D {
	editor := material.Editor(th, &vc.inpVal.Editor, "Input incr/decr value")

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,
		layout.Flexed(1, func(gtx C) D {
			vc.inpVal.Editor.SingleLine = true
			vc.inpVal.Editor.Alignment = text.Middle
			editor.Font.Weight = text.Bold
			editor.TextSize = unit.Sp(20)
			border := widget.Border{
				Color:        colours["grey"],
				CornerRadius: unit.Dp(5),
				Width:        unit.Px(3),
			}
			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Sp(8)).Layout(
					gtx,
					editor.Layout,
				)
			})
		}),

		layout.Rigid(
			layout.Spacer{
				Width: defaultMargin,
			}.Layout,
		),

		layout.Flexed(1, func(gtx C) D {
			btn := material.Button(th, &vc.changeVal, "Change val to incr/decr")
			btn.Background = colours["blue"]
			if vc.changeVal.Clicked() {
				inpVal := vc.inpVal.Text()
				inpVal = strings.TrimSpace(inpVal)
				intVal, _ := strconv.ParseInt(inpVal, 10, 8)
				count = int(intVal)
			}
			return btn.Layout(gtx)
		}),
	)
}

func (ui *UI) Layout(gtx C) D {
	inset := layout.UniformInset(defaultMargin)
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,
		layout.Flexed(0.1, func(gtx C) D {
			return inset.Layout(gtx, func(gtx C) D {
				return ui.valueChanger.Layout(ui.theme, gtx)
			})
		}),
		layout.Flexed(1, func(gtx C) D {
			return inset.Layout(gtx, func(gtx C) D {
				return ui.counter.Layout(ui.theme, gtx)
			})
		}),
	)
}
