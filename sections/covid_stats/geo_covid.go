package covid_stats

import (
	"fmt"
	"gioui-experiment/globals"
	"gioui-experiment/sections"
	"gioui-experiment/sections/covid_stats/data"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/outlay"
	"image"
	"image/color"
	"log"
	"time"
)

// twoDaysPrior - the default start date -  represents 2 days before today. format: MM-DD-YYYY
var twoDaysPrior = time.Now().Local().AddDate(0, 0, -2).Format("1-02-2006")

type (
	C = layout.Context
	D = layout.Dimensions

	Section struct {
		*sections.Router
		th *material.Theme

		Display

		DisableCPBtn widget.Clickable
		isCPDisabled bool
	}

	Display struct {
		Api data.CovidStats
		UiTable
	}

	UiTable struct {
		yList widget.List
		cells []TableCell
		outlay.Table
	}

	TableCell struct {
		btn     widget.Clickable
		clicked bool
	}
)

func CovidTable() *UiTable {
	return &UiTable{
		Table: outlay.Table{
			CellSize: func(m unit.Metric, x, y int) image.Point {
				return image.Pt(m.Px(unit.Dp(100)), m.Px(unit.Dp(75)))
			},
		},
	}
}

func (s *Section) LayoutView(gtx C, th *material.Theme) D {
	s.yList.Axis = layout.Vertical

	err := data.InitDayData(twoDaysPrior)
	if err != nil {
		log.Fatalln(err)
	}
	return material.List(th, &s.yList).Layout(gtx, 1, func(gtx C, _ int) D {
		return CovidTable().Layout(gtx, th)
	})
}

func (s *Section) LayoutController(gtx C, th *material.Theme) D {
	return D{}
}

func New(router *sections.Router) *Section {
	return &Section{
		Router: router,
		th:     material.NewTheme(gofont.Collection()),
	}
}

func (s *Section) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Tag: &s.DisableCPBtn,
			},
			Layout: func(gtx C, bg, fg color.NRGBA) D {
				var lbl string
				if s.DisableCPBtn.Clicked() {
					s.isCPDisabled = !s.isCPDisabled
				}
				if !s.isCPDisabled {
					lbl = "Disable CP"
				} else {
					lbl = "Enable CP"
				}
				return material.Button(s.th, &s.DisableCPBtn, lbl).Layout(gtx)
			},
		},
	}
}

func (s *Section) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (s *Section) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Covid-19 Statistics",
	}
}

func (s *Section) IsCPDisabled() bool {
	return s.isCPDisabled
}

func (uit *UiTable) Layout(gtx C, th *material.Theme) D {
	xn := 1
	yn := len(data.CachedDays[twoDaysPrior])
	uit.cells = growCells(uit.cells, xn*yn)

	return uit.Table.Layout(gtx, xn, yn, func(gtx C, x, y int) D {
		c := &uit.cells[x+y]
		return c.Layout(gtx, th, x, y)
	})
}

func (tc *TableCell) Layout(gtx C, th *material.Theme, x, y int) D {
	var txt string
	if y == 0 {
		txt = fmt.Sprintf("Location")
	} else {
		txt = fmt.Sprintf(data.CachedDays[twoDaysPrior][y-1].CombinedKey)
	}

	macro := op.Record(gtx.Ops)
	dims := layout.Center.Layout(gtx, material.Body1(th, txt).Layout)
	call := macro.Stop()

	defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, globals.Colours[globals.ANTIQUE_WHITE])

	call.Add(gtx.Ops)
	return dims
}

func growCells(cells []TableCell, n int) []TableCell {
	if len(cells) < n {
		cells = append(cells, make([]TableCell, n-len(cells))...)
	}
	return cells[:n]
}
