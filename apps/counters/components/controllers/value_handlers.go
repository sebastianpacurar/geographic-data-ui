package controllers

import (
	"gioui-experiment/apps/counters/components/data"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"strconv"
	"strings"
	"unicode/utf8"
)

type ValueHandler struct {
	startVal, unitVal       component.TextField
	changeStart, changeUnit widget.Clickable
}

func (vh *ValueHandler) InitTextFields() {
	vh.startVal.SingleLine = true
	vh.unitVal.SingleLine = true
	vh.startVal.Alignment = layout.Alignment(text.Start)
	vh.unitVal.Alignment = layout.Alignment(text.Start)
}

func (vh *ValueHandler) Layout(th *material.Theme, gtx C) D {
	cv := data.CounterVals
	eStart := material.Editor(th, &vh.startVal.Editor, "0")
	eUnit := material.Editor(th, &vh.unitVal.Editor, "1")

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,

		layout.Rigid(func(gtx C) D {
			if !isFieldNumeric(vh.startVal) {
				gtx = gtx.Disabled()
			}
			btn := material.Button(th, &vh.changeStart, "Start From")
			btn.Background = g.Colours["blue"]
			btn.Color = g.Colours["white"]
			for range vh.changeStart.Clicks() {
				vh.handleStartBtn(vh.startVal, cv)
			}
			return btn.Layout(gtx)
		}),

		g.SpacerX,

		layout.Flexed(1, func(gtx C) D {
			eStart.TextSize = unit.Sp(18)
			eStart.HintColor = g.Colours["dark-slate-grey"]
			border := g.DefaultBorder
			vh.validateTextField(th, vh.startVal, eStart, &border)
			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(
					gtx,
					eStart.Layout,
				)
			})
		}),

		g.SpacerX,

		layout.Rigid(func(gtx C) D {
			btn := material.Button(th, &vh.changeUnit, "Set Unit To")
			btn.Background = g.Colours["blue"]
			if !isFieldNumeric(vh.unitVal) {
				gtx = gtx.Disabled()
			}

			for range vh.changeUnit.Clicks() {
				vh.handleUnitBtn(vh.unitVal, cv)
			}
			return btn.Layout(gtx)
		}),

		g.SpacerX,

		layout.Flexed(1, func(gtx C) D {
			eUnit.TextSize = unit.Sp(18)
			eUnit.HintColor = g.Colours["dark-slate-grey"]
			border := g.DefaultBorder
			vh.validateTextField(th, vh.unitVal, eUnit, &border)
			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(
					gtx,
					eUnit.Layout,
				)
			})
		}),
	)
}

// isFieldNumeric - verifies if the string can be converted to integer.
// It iterates through the input and verifies if each rune digit casted to string,
// can be converted to integer. Exception from the rule is the minus sign as first char.
func isFieldNumeric(e component.TextField) bool {
	if e.Len() > 0 {
		for i := range e.Text() {
			if i == 0 && e.Text()[0] == '-' {
				continue
			}
			_, err := strconv.Atoi(string(e.Text()[i]))
			if err != nil {
				return false
			}
		}
	}
	return true
}

// trimInput - removes the last rune from the string iteratively. Used in case
// the input char count passes 17 chars
func trimInput(e component.TextField, count int) {
	str := e.Text()
	for i := 0; i < count; i++ {
		_, _ = utf8.DecodeLastRuneInString(str)
	}
	e.SetText(str[:len(str)-count])
}

// 1) In case the length of the input is >= 18, then keep trimming chars
// until the length is back to 17
// 2) If there are non-numeric characters, change colors to red
// 3) When Focused change border color a bit
func (vh *ValueHandler) validateTextField(th *material.Theme, e component.TextField, eStyle material.EditorStyle, b *widget.Border) {
	switch {
	case e.Len() >= 18:
		trimInput(e, e.Len()-18)
	case !isFieldNumeric(e):
		b.Color = g.Colours["red"]
		b.Width = unit.Px(5)
		eStyle.Color = g.Colours["dark-red"]
	case e.Focused():
		b.Color = th.Palette.ContrastBg
		b.Width = unit.Px(3)
	}
}

func (vh *ValueHandler) handleStartBtn(e component.TextField, cv *data.CurrentValues) {
	val := strings.TrimSpace(e.Text())
	switch cv.GetActiveSequence() {
	case data.PRIMES:
		numVal, _ := strconv.ParseUint(val, 10, 64)
		cv.PCurrIndex = int(numVal)
		cv.PCount = cv.PCache[cv.PCurrIndex]
	case data.FIBS:
		numVal, _ := strconv.ParseUint(val, 10, 64)
		cv.FCurrIndex = int(numVal)
		cv.FCount = cv.FCache[cv.FCurrIndex]
	case data.NATURALS:
		numVal, _ := strconv.ParseUint(val, 10, 64)
		cv.NCount = numVal
	case data.WHOLES:
		numVal, _ := strconv.ParseInt(val, 10, 64)
		cv.WCount = numVal
	}
}

func (vh *ValueHandler) handleUnitBtn(e component.TextField, cv *data.CurrentValues) {
	val := strings.TrimSpace(e.Text())
	switch cv.GetActiveSequence() {
	case data.PRIMES:
		numVal, _ := strconv.ParseUint(val, 10, 64)
		cv.PCurrIndex = int(numVal)
		cv.PCount = cv.PCache[cv.PCurrIndex]
	case data.FIBS:
		numVal, _ := strconv.ParseUint(val, 10, 64)
		cv.FCurrIndex = int(numVal)
		cv.FCount = cv.FCache[cv.FCurrIndex]
	case data.NATURALS:
		numVal, _ := strconv.ParseUint(val, 10, 64)
		cv.NCount = numVal
	case data.WHOLES:
		numVal, _ := strconv.ParseInt(val, 10, 64)
		cv.WCount = numVal
	}
}
