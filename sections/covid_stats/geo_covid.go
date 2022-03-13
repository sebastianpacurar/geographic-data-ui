package covid_stats

import (
	"fmt"
	"gioui-experiment/sections"
	"gioui-experiment/sections/covid_stats/data"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image/color"
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
	}
)

func (s *Section) LayoutView(gtx C, th *material.Theme) D {
	err := data.InitDayData(twoDaysPrior)
	fmt.Println(data.CachedDays[twoDaysPrior])
	if err != nil {
		return D{}
	}
	return D{}
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
		Name: "CovidStats-19 Statistics",
	}
}

func (s *Section) IsCPDisabled() bool {
	return s.isCPDisabled
}
