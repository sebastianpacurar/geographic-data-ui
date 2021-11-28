package controllers

import (
	"gioui-experiment/apps/counters/components/data"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type Status struct {
	primes   data.Primes
	fibs     data.Fibonacci
	naturals data.Naturals
	integers data.Integers

	primesState   component.DiscloserState
	fibsState     component.DiscloserState
	naturalsState component.DiscloserState
	integersState component.DiscloserState
}

func (s *Status) Layout(gtx C, th *material.Theme) D {
	return material.Body2(th, "In Progress").Layout(gtx)
}
