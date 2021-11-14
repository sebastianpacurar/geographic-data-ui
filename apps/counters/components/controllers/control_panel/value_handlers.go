package control_panel

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
)

type (
	ValueHandler struct {
		startFrom
		skipBy
	}

	startFrom struct {
		textField component.TextField
		btn       widget.Clickable
		isValid   bool
	}

	skipBy struct {
		textField component.TextField
		btn       widget.Clickable
		isValid   bool
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

	// lay startRow = horizontal layout for startFromField - startFromBtn
	startRow := layout.Rigid(func(gtx C) D {
		return layout.Flex{}.Layout(gtx, startFromField, startFromBtn)
	})

	// lay skipRow = horizontal layout for skipByField - skipByBtn
	skipRow := layout.Rigid(func(gtx C) D {
		return layout.Flex{}.Layout(gtx, skipByField, skipByBtn)
	})

	// lay out startRow and skipRow vertically
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		startRow, skipRow,
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
	case data.PRIMES, data.FIBS:
		cv.Index = int(numVal) - 1
	case data.NATURALS, data.INTEGERS:
		cv.Displayed = numVal
	}
}

func (vh *ValueHandler) handleSkipBtn(cv *data.Generator) {
	val := strings.TrimSpace(vh.skipBy.textField.Text())
	numVal, _ := strconv.ParseUint(val, 10, 64)
	cv.Step = numVal
}
