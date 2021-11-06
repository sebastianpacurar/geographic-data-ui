package controllers

import (
	"fmt"
	"gioui-experiment/apps/counters/components/data"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"strconv"
	"strings"
	"unicode"
)

type (
	ValueHandler struct {
		startFrom
		skipBy
		resetTo
	}

	startFrom struct {
		textField component.TextField
		btn       widget.Clickable
	}

	skipBy struct {
		textField component.TextField
		btn       widget.Clickable
	}

	resetTo struct {
		textField component.TextField
		btn       widget.Clickable
	}
)

func (vh *ValueHandler) Layout(gtx C, th *material.Theme) D {
	cv := data.CurrVals
	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Flexed(1, func(gtx C) D {
					return vh.InputBox(gtx, th, &vh.startFrom.textField, cv, "start")
				}),
				layout.Flexed(1, func(gtx C) D {
					for range vh.startFrom.btn.Clicks() {
						vh.handleStartBtn(cv)
					}
					return g.Inset.Layout(gtx, func(C) D {
						return material.Button(th, &vh.startFrom.btn, "set start").Layout(gtx)
					})
				}),
			)
		}),
		layout.Rigid(func(gtx C) D {
			return layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Flexed(1, func(gtx C) D {
					return vh.InputBox(gtx, th, &vh.skipBy.textField, cv, "skip")
				}),
				layout.Flexed(1, func(gtx C) D {
					for range vh.skipBy.btn.Clicks() {
						vh.handleSkipBtn(cv)
					}
					return g.Inset.Layout(gtx, func(C) D {
						return material.Button(th, &vh.skipBy.btn, "set step").Layout(gtx)
					})
				}),
			)
		}))
}

func (vh *ValueHandler) InputBox(gtx C, th *material.Theme, e *component.TextField, cv *data.Generator, context string) D {
	if !isNumeric(e.Text()) {
		e.SetError("only digits allowed")
	} else {
		e.ClearError()
	}
	seq := cv.GetActiveSequence()
	var placeholder string
	switch seq {
	case data.PRIMES, data.FIBS:
		switch context {
		case "start":
			placeholder = fmt.Sprintf("n-th %s", seq[:len(seq)-1])
			e.CharLimit = 5
		case "reset":
			placeholder = "set reset to n"
			e.CharLimit = 10
		case "skip":
			placeholder = "set skip to n"
			e.CharLimit = 5
		}
	case data.NATURALS, data.WHOLES:
		switch context {
		case "start":
			placeholder = "start from n"
			e.CharLimit = 15
		case "skip":
			placeholder = "set skip to n"
			e.CharLimit = 15
		}
	}
	if e.Len() > int(e.CharLimit) {
		e.SetError("limit exceeded")
	} else {
		e.ClearError()
	}
	e.SingleLine = true
	e.Alignment = layout.Alignment(text.Start)
	return e.Layout(gtx, th, placeholder)
}

func isNumeric(val string) bool {
	for _, r := range val {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func (vh *ValueHandler) handleStartBtn(cv *data.Generator) {
	val := strings.TrimSpace(vh.startFrom.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES, data.FIBS:
		cv.Index = int(numVal) - 1
	case data.NATURALS, data.WHOLES:
		cv.Displayed = numVal
	}
}

func (vh *ValueHandler) handleSkipBtn(cv *data.Generator) {
	val := strings.TrimSpace(vh.skipBy.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	cv.Step = numVal
}
