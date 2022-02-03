package playground

import (
	"gioui-experiment/apps"
	"gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/font/gofont"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
	"image/color"
	"time"
)

var (
	startTime       = time.Now()
	duration        = 2 * time.Second
	controllerInset = layout.Inset{
		Top:    unit.Dp(10),
		Right:  unit.Dp(25),
		Bottom: unit.Dp(10),
	}
	sqColors = map[string]color.NRGBA{
		"left":     globals.Colours[colours.LIGHT_SEA_GREEN],
		"right":    globals.Colours[colours.FLAME_RED],
		"middle":   globals.Colours[colours.BLACK],
		"dragged":  globals.Colours[colours.GREY],
		"inactive": globals.Colours[colours.ELECTRIC_BLUE],
	}
	bg color.NRGBA
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		th *material.Theme
		ControlPanel
		*apps.Router

		DisableCPBtn widget.Clickable
		isCPDisabled bool

		Draw
	}
	Draw struct {
		active bool
		sq     Square
	}

	Square struct {
		pos  int
		drag gesture.Drag
	}
)

func New(router *apps.Router) *Application {
	return &Application{
		Router: router,
		th:     material.NewTheme(gofont.Collection()),
	}
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Tag: &app.DisableCPBtn,
			},
			Layout: func(gtx C, bg, fg color.NRGBA) D {
				var (
					lbl string
				)
				if app.DisableCPBtn.Clicked() {
					app.isCPDisabled = !app.isCPDisabled
				}

				if !app.isCPDisabled {
					lbl = "Disable CP"
				} else {
					lbl = "Enable CP"
				}
				return material.Button(app.th, &app.DisableCPBtn, lbl).Layout(gtx)
			},
		},
	}
}

func (app *Application) Overflow() []component.OverflowAction {
	return []component.OverflowAction{
		{Name: "Close Current Instance - dummy action"},
	}
}

func (app *Application) NavItem() component.NavItem {
	return component.NavItem{
		Name: "playground",
	}
}

func (app *Application) IsCPDisabled() bool {
	return app.isCPDisabled
}

func (app *Application) LayoutView(gtx C, th *material.Theme) D {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return app.Draw.Layout(gtx, th)
		}))
}

func (app *Application) LayoutController(gtx C, th *material.Theme) D {
	return app.ControlPanel.Layout(gtx, th)
}

func (d *Draw) Layout(gtx C, th *material.Theme) D {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
			return globals.RColoredArea(gtx, size, 10, globals.Colours[colours.WHITE])
		}),
		layout.Stacked(func(gtx C) D {
			return layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical, WeightSum: 3}.Layout(gtx,
					layout.Flexed(1, func(gtx C) D {
						return d.drawInteractiveRect(gtx)
					}),

					layout.Rigid(func(gtx C) D {
						return layout.Inset{Top: unit.Dp(5), Bottom: unit.Dp(5)}.Layout(gtx, func(gtx C) D {
							return d.drawDelimiter(gtx)
						})
					}),

					layout.Flexed(1, func(gtx C) D {
						return d.fillGradually(gtx, gtx.Now)
					}),

					layout.Rigid(func(gtx C) D {
						return layout.Inset{Top: unit.Dp(5), Bottom: unit.Dp(5)}.Layout(gtx, func(gtx C) D {
							return d.drawDelimiter(gtx)
						})
					}),

					// does nothing
					layout.Flexed(1, func(gtx C) D {
						return d.sq.drawDraggableSquare(gtx)
					}),
				)
			})
		}))
}

func (sq *Square) drawDraggableSquare(gtx C) D {
	dims := image.Point{X: 100, Y: 100}

	var de *pointer.Event
	for _, e := range sq.drag.Events(gtx.Metric, gtx, gesture.Axis(layout.Horizontal)) {
		if e.Type == pointer.Drag {
			de = &e
		}
	}
	if de != nil {
		xy := de.Position.X
		sq.pos += int(xy)
	}
	sq.drag.Add(gtx.Ops)
	rect := clip.Rect{Max: dims}
	paint.FillShape(gtx.Ops, sqColors["inactive"], rect.Op())
	return D{Size: rect.Max}
}

func (d *Draw) drawInteractiveRect(gtx C) D {
	for _, ev := range gtx.Events(&d.active) {
		switch ev := ev.(type) {
		case pointer.Event:
			switch ev.Type {
			case pointer.Press:
				d.active = true
				switch ev.Buttons {
				case pointer.ButtonPrimary:
					bg = sqColors["left"]
				case pointer.ButtonSecondary:
					bg = sqColors["right"]
				case pointer.ButtonTertiary:
					bg = sqColors["middle"]
				}
			case pointer.Drag:
				bg = sqColors["dragged"]

			case pointer.Release, pointer.Cancel:
				d.active = false
				bg = sqColors["inactive"]
			}
		}
	}

	area := clip.Rect(image.Rect(0, 0, 100, 100)).Push(gtx.Ops)
	pointer.CursorNameOp{Name: pointer.CursorGrab}.Add(gtx.Ops)
	pointer.InputOp{
		Tag:   &d.active,
		Types: pointer.Press | pointer.Release | pointer.Cancel | pointer.Drag,
		Grab:  true,
	}.Add(gtx.Ops)
	area.Pop()

	rect := clip.Rect{Max: image.Pt(100, 100)}
	paint.FillShape(gtx.Ops, bg, rect.Op())

	return D{Size: rect.Max}
}

func (d *Draw) fillGradually(gtx C, now time.Time) D {
	elapsed := now.Sub(startTime)
	progress := elapsed.Seconds() / duration.Seconds()
	if progress < 1 {
		op.InvalidateOp{}.Add(gtx.Ops)
	} else {
		progress = 0
	}

	width := float32(gtx.Constraints.Max.X) * float32(progress)
	height := float32(gtx.Constraints.Max.Y) * float32(progress)
	rect := clip.Rect{Max: image.Pt(int(width), int(height))}

	paint.FillShape(gtx.Ops, globals.Colours[colours.AERO_BLUE], rect.Op())
	return D{Size: rect.Max}
}

func (d *Draw) drawDelimiter(gtx C) D {
	rect := clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, 10)}
	paint.FillShape(gtx.Ops, globals.Colours[colours.ANTIQUE_WHITE], rect.Op())
	return D{Size: rect.Max}
}

func (d *Draw) dragArea(gtx C) D {
	return D{}
}
