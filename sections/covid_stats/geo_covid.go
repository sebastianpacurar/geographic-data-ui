package covid_stats

import (
	"gioui-experiment/sections"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Section struct {
		*sections.Router
		th *material.Theme

		DisableCPBtn widget.Clickable
		isCPDisabled bool
	}
)

func (s *Section) LayoutView(gtx C, th *material.Theme) D {
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
		Name: "Covid-19 Statistics",
	}
}

func (s *Section) IsCPDisabled() bool {
	return s.isCPDisabled
}
