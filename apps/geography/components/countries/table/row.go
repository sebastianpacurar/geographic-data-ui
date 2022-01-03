package table

import (
	"fmt"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
)

type Row struct {
	Click        widget.Clickable
	Name         string
	OfficialName string
	Capital      []string
	Independent  bool
	Status       string
	UNMember     bool
	Cca2         string
	Cca3         string
	Ccn3         string
	Area         float64
	Population   int32
	Active       bool
	Selected     bool
	IsCPViewed   bool
	IsViewed     bool
}

func (r *Row) LayRow(gtx C, th *material.Theme) D {
	border := widget.Border{
		Color: g.Colours[colors.GREY],
		Width: unit.Px(1),
	}

	return material.Clickable(gtx, &r.Click, func(gtx C) D {

		rowColor := g.Colours[colors.ANTIQUE_WHITE]

		if r.Selected {
			rowColor = g.Colours[colors.AERO_BLUE]
		}

		if r.Click.Hovered() {
			if r.Selected {
				rowColor = g.Colours[colors.LIGHT_SALMON]
			} else {
				rowColor = g.Colours[colors.NYANZA]
			}
		}

		row := layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(450, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(func(gtx C) D {
								return material.Body1(th, r.Name).Layout(gtx)
							}))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(550, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(func(gtx C) D {
								return material.Body1(th, r.OfficialName).Layout(gtx)
							}))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(200, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(func(gtx C) D {
								Capital := "N/A"
								if len(r.Capital) > 0 {
									Capital = r.Capital[0]
								}
								return material.Body1(th, Capital).Layout(gtx)
							}))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(func(gtx C) D {
								Independent := "No"
								if r.Independent {
									Independent = "Yes"
								}
								return material.Body1(th, Independent).Layout(gtx)
							}))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(175, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(func(gtx C) D {
								return material.Body1(th, r.Status).Layout(gtx)
							}))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(func(gtx C) D {
								UNMember := "No"
								if r.UNMember {
									UNMember = "Yes"
								}
								return material.Body1(th, UNMember).Layout(gtx)
							}))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(material.Body1(th, r.Cca2).Layout))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(material.Body1(th, r.Cca3).Layout))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(50, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(material.Body1(th, r.Ccn3).Layout))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(100, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(material.Body1(th, fmt.Sprintf("%.0f", r.Area)).Layout))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return border.Layout(gtx, func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(2),
						Bottom: unit.Dp(2),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						return layout.Stack{Alignment: layout.Center}.Layout(gtx,

							layout.Expanded(func(gtx C) D {
								return g.ColoredArea(gtx, image.Pt(100, gtx.Constraints.Min.Y), rowColor)
							}),

							layout.Stacked(material.Body1(th, fmt.Sprintf("%d", int(r.Population))).Layout))
					})
				})
			}),
		)
		return row
	})
}
