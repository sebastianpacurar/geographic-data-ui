package controllers

import (
	"fmt"
	"gioui-experiment/apps/playground/components/data"
	g "gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"strconv"
	"strings"
)

type (
	ValueHandler struct {
		startFrom
		skipBy
		stopAt
	}

	startFrom struct {
		textField component.TextField
		btn       widget.Clickable
	}

	skipBy struct {
		textField component.TextField
		btn       widget.Clickable
	}

	stopAt struct {
		textField component.TextField
		btn       widget.Clickable
	}
)

func (vh *ValueHandler) Layout(gtx C, th *material.Theme) D {
	cv := data.CurrVals
	// start field
	startFromField := layout.Flexed(1, func(gtx C) D {
		return vh.InputBox(gtx, th, &vh.startFrom.textField, cv, "start")
	})
	// start button
	startFromBtn := layout.Flexed(1, func(gtx C) D {
		field := vh.startFrom.textField
		if field.IsErrored() || field.Len() == 0 {
			gtx = gtx.Disabled()
		}
		for range vh.startFrom.btn.Clicks() {
			vh.handleStartBtn(cv)
		}
		return g.Inset.Layout(gtx, func(C) D {
			return material.Button(th, &vh.startFrom.btn, "set start").Layout(gtx)
		})
	})
	// skip (step) field
	skipByField := layout.Flexed(1, func(gtx C) D {
		return vh.InputBox(gtx, th, &vh.skipBy.textField, cv, "skip")
	})
	// skip (step) button
	skipByBtn := layout.Flexed(1, func(gtx C) D {
		field := vh.skipBy.textField
		if field.IsErrored() || field.Len() == 0 {
			gtx = gtx.Disabled()
		}
		for range vh.skipBy.btn.Clicks() {
			vh.handleSkipBtn(cv)
		}
		return g.Inset.Layout(gtx, func(C) D {
			return material.Button(th, &vh.skipBy.btn, "set step").Layout(gtx)
		})
	})

	//TODO: temporary on hold
	// stop field
	stopAtField := layout.Flexed(1, func(gtx C) D {
		gtx = gtx.Disabled()
		return vh.InputBox(gtx, th, &vh.stopAt.textField, cv, "stop")
	})
	//TODO: temporary on hold
	// stop button
	stopAtBtn := layout.Flexed(1, func(gtx C) D {
		gtx = gtx.Disabled()
		field := vh.stopAt.textField
		if field.IsErrored() || field.Len() == 0 {
			gtx = gtx.Disabled()
		}
		for range vh.stopAt.btn.Clicks() {
			vh.handleStopBtn(cv)
		}
		return g.Inset.Layout(gtx, func(C) D {
			return material.Button(th, &vh.stopAt.btn, "on hold").Layout(gtx)
		})
	})

	startRow := layout.Rigid(func(gtx C) D {
		return layout.Flex{}.Layout(gtx, startFromField, startFromBtn)
	})

	skipRow := layout.Rigid(func(gtx C) D {
		return layout.Flex{}.Layout(gtx, skipByField, skipByBtn)
	})

	stopRow := layout.Rigid(func(gtx C) D {
		return layout.Flex{}.Layout(gtx, stopAtField, stopAtBtn)
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		startRow, skipRow, stopRow,
	)
}

func (vh *ValueHandler) InputBox(gtx C, th *material.Theme, e *component.TextField, cv *data.Generator, context string) D {
	seq := cv.GetActiveSequence()
	var placeholder string
	switch seq {
	case data.PRIMES, data.FIBS:
		switch context {
		case "start":
			placeholder = fmt.Sprintf("n-th %s", seq[:len(seq)-1])
			e.CharLimit = 5
		case "skip":
			placeholder = "set step by n"
			e.CharLimit = 5
		}
	case data.NATURALS, data.INTEGERS:
		switch context {
		case "start":
			placeholder = "start from n"
			e.CharLimit = 15
		case "skip":
			placeholder = "set step by n"
			e.CharLimit = 15
		}
	}
	if e.Len() > int(e.CharLimit) {
		e.SetError("limit exceeded")
	} else if !isNumeric(e.Text()) && e.Len() > 0 {
		e.SetError("only digits")
	} else {
		e.ClearError()
	}

	e.SingleLine = true
	e.Alignment = layout.Alignment(text.Start)
	return e.Layout(gtx, th, placeholder)
}

func isNumeric(val string) bool {
	if _, err := strconv.ParseInt(val, 10, 64); err != nil {
		return false
	} else {
		return true
	}
}

func (vh *ValueHandler) handleStartBtn(cv *data.Generator) {
	val := strings.TrimSpace(vh.startFrom.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		cv.Primes.Index = int(numVal) - 1
	case data.FIBS:
		cv.Fibonacci.Index = int(numVal) - 1
	case data.NATURALS:
		cv.Naturals.Displayed = numVal
	case data.INTEGERS:
		cv.Integers.Displayed = numVal
	}
}

func (vh *ValueHandler) handleSkipBtn(cv *data.Generator) {
	val := strings.TrimSpace(vh.skipBy.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		cv.Primes.Step = numVal
	case data.FIBS:
		cv.Fibonacci.Step = numVal
	case data.NATURALS:
		cv.Naturals.Step = numVal
	case data.INTEGERS:
		cv.Integers.Step = numVal
	}
}

func (vh *ValueHandler) handleStopBtn(cv *data.Generator) {
	val := strings.TrimSpace(vh.stopAt.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	seq := cv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		cv.Primes.Stop = numVal
	case data.FIBS:
		cv.Fibonacci.Stop = numVal
	}
}
