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
		toggle    widget.Bool
	}

	skipBy struct {
		textField component.TextField
		toggle    widget.Bool
	}

	resetTo struct {
		textField component.TextField
		toggle    widget.Bool
	}
)

func (vh *ValueHandler) Layout(th *material.Theme, gtx C) D {
	cv := data.CounterVals
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,
		layout.Flexed(1, func(gtx C) D {
			return vh.InputBox(gtx, th, &vh.startFrom.textField, cv, "start")
		}),

		g.SpacerX,

		layout.Flexed(1, func(gtx C) D {
			return g.Inset.Layout(gtx, func(C) D {
				vh.handleSkipToggle(cv)
				return material.CheckBox(th, &vh.startFrom.toggle, "").Layout(gtx)
			})
		}),

		g.SpacerX,

		layout.Flexed(1, func(gtx C) D {
			return vh.InputBox(gtx, th, &vh.skipBy.textField, cv, "skip")
		}),

		g.SpacerX,

		layout.Flexed(1, func(gtx C) D {
			return g.Inset.Layout(gtx, func(C) D {
				vh.handleStartToggle(cv)
				return material.CheckBox(th, &vh.skipBy.toggle, "").Layout(gtx)
			})
		}),
	)
}

func (vh *ValueHandler) InputBox(gtx C, th *material.Theme, e *component.TextField, cv *data.CurrentValues, context string) D {
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
	default:
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

func (vh *ValueHandler) handleStartToggle(cv *data.CurrentValues) {
	if vh.startFrom.toggle.Changed() {
		if vh.startFrom.toggle.Value {
			val := strings.TrimSpace(vh.startFrom.textField.Text())
			numVal, _ := strconv.ParseUint(val, 10, 64)
			cv.Index = int(numVal) - 1
			cv.Start = uint64(cv.Index)
			cv.Displayed = numVal
		} else {
			cv.Index = 0
			cv.Start = data.ONE
			cv.Displayed = data.ONE
		}
	}
}

func (vh *ValueHandler) handleSkipToggle(cv *data.CurrentValues) {
	if vh.skipBy.toggle.Changed() {
		if vh.skipBy.toggle.Value {
			val := strings.TrimSpace(vh.skipBy.textField.Text())
			numVal, _ := strconv.ParseUint(val, 10, 64)
			cv.Step = numVal
		} else {
			cv.Step = data.ONE
		}
	}
}
