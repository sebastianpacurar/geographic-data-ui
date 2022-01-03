package controllers

import (
	"fmt"
	"gioui-experiment/apps/playground/data"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
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

// Layout TODO: this gets way too overly complicated. To Be Simplified!
func (vh *ValueHandler) Layout(gtx C, th *material.Theme) D {
	pgv := data.PgVals
	// start field
	startFromField := layout.Flexed(1, func(gtx C) D {
		return vh.InputBox(gtx, th, &vh.startFrom.textField, pgv, "start")
	})
	// start button
	startFromBtn := layout.Flexed(1, func(gtx C) D {
		field := vh.startFrom.textField
		if field.IsErrored() || field.Len() == 0 {
			gtx = gtx.Disabled()
		}
		for range vh.startFrom.btn.Clicks() {
			vh.handleStartBtn(pgv)
		}
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(C) D {
			return material.Button(th, &vh.startFrom.btn, "set start").Layout(gtx)
		})
	})
	// skip (step) field
	skipByField := layout.Flexed(1, func(gtx C) D {
		return vh.InputBox(gtx, th, &vh.skipBy.textField, pgv, "skip")
	})
	// skip (step) button
	skipByBtn := layout.Flexed(1, func(gtx C) D {
		field := vh.skipBy.textField
		if field.IsErrored() || field.Len() == 0 {
			gtx = gtx.Disabled()
		}
		for range vh.skipBy.btn.Clicks() {
			vh.handleSkipBtn(pgv)
		}
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(C) D {
			return material.Button(th, &vh.skipBy.btn, "set step").Layout(gtx)
		})
	})

	//TODO: temporary on hold
	// stop field
	stopAtField := layout.Flexed(1, func(gtx C) D {
		gtx = gtx.Disabled()
		return vh.InputBox(gtx, th, &vh.stopAt.textField, pgv, "stop")
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
			vh.handleStopBtn(pgv)
		}
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(C) D {
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

func (vh *ValueHandler) InputBox(gtx C, th *material.Theme, e *component.TextField, pgv *data.Generator, context string) D {
	seq := pgv.GetActiveSequence()
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

func (vh *ValueHandler) handleStartBtn(pgv *data.Generator) {
	val := strings.TrimSpace(vh.startFrom.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	seq := pgv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		pgv.Primes.Index = int(numVal) - 1
	case data.FIBS:
		pgv.Fibonacci.Index = int(numVal) - 1
	case data.NATURALS:
		pgv.Naturals.Displayed = numVal
	case data.INTEGERS:
		pgv.Integers.Displayed = numVal
	}
}

func (vh *ValueHandler) handleSkipBtn(pgv *data.Generator) {
	val := strings.TrimSpace(vh.skipBy.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	seq := pgv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		pgv.Primes.Step = numVal
	case data.FIBS:
		pgv.Fibonacci.Step = numVal
	case data.NATURALS:
		pgv.Naturals.Step = numVal
	case data.INTEGERS:
		pgv.Integers.Step = numVal
	}
}

func (vh *ValueHandler) handleStopBtn(pgv *data.Generator) {
	val := strings.TrimSpace(vh.stopAt.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	seq := pgv.GetActiveSequence()
	switch seq {
	case data.PRIMES:
		pgv.Primes.Stop = numVal
	case data.FIBS:
		pgv.Fibonacci.Stop = numVal
	}
}
