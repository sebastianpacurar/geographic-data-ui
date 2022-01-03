package playground

import (
	"gioui-experiment/apps"
	"gioui-experiment/apps/playground/data/counter"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"strconv"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Application struct {
		//dockBtn widget.Clickable
		//btn     material.IconButtonStyle
		//icon    *widget.Icon
		th *material.Theme
		ControlPanel
		*apps.Router
	}
)

func New(router *apps.Router) *Application {
	return &Application{
		Router: router,
	}
}

func (app *Application) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		// TODO: think what to do with this only on PG
		//{
		//	OverflowAction: component.OverflowAction{
		//		Tag: &app.dockBtn,
		//	},
		//	Layout: func(gtx C, bg, fg color.NRGBA) D {
		//		for range app.dockBtn.Clicks() {
		//			app.NonModalDrawer = !app.NonModalDrawer
		//		}
		//		if app.NonModalDrawer {
		//			app.icon = g.LockCLosedIcon
		//			app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
		//			app.btn.Background = bg
		//			app.btn.Color = g.Colours[colors.DARK_RED]
		//			app.btn.Size = unit.Dp(24)
		//		} else {
		//			app.icon = g.LockOpenedIcon
		//			app.btn = component.SimpleIconButton(bg, fg, &app.dockBtn, app.icon)
		//			app.btn.Background = bg
		//			app.btn.Color = g.Colours[colors.SEA_GREEN]
		//			app.btn.Size = unit.Dp(24)
		//		}
		//		return app.btn.Layout(gtx)
		//	},
		//},
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

func (app *Application) LayoutView(gtx C, th *material.Theme) D {
	pgv := counter.PgVals

	//TODO find a better way and location to handle the Cache population
	if len(pgv.Cache[counter.PRIMES]) == 0 {
		pgv.GenPrimes(counter.PLIMIT)
	}
	if len(pgv.Cache[counter.FIBS]) == 0 {
		pgv.GenFibs(counter.FLIMIT)
	}
	seq := pgv.GetActiveSequence()

	/// DISPLAYED NUMBER
	return layout.Inset{
		Top:    unit.Dp(10),
		Right:  unit.Dp(50),
		Bottom: unit.Dp(20),
		Left:   unit.Dp(50),
	}.Layout(gtx, func(gtx C) D {
		var val string
		switch seq {
		case counter.PRIMES:
			val = strconv.FormatUint(pgv.Cache[seq][pgv.Primes.Index], 10)
		case counter.FIBS:
			val = strconv.FormatUint(pgv.Cache[seq][pgv.Fibonacci.Index], 10)
		case counter.NATURALS:
			val = strconv.FormatUint(pgv.Naturals.Displayed, 10)
		case counter.INTEGERS:
			val = strconv.FormatUint(pgv.Integers.Displayed, 10)
		}
		return material.H5(th, val).Layout(gtx)
	})
}

func (app *Application) LayoutController(gtx C, th *material.Theme) D {
	return app.ControlPanel.Layout(gtx, th)
}
