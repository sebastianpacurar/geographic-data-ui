package controllers

import (
	"gioui-experiment/apps/playground/components/data/counter"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	Sequence struct {
		radioBtns widget.Enum
	}
)

func (s *Sequence) Layout(gtx C, th *material.Theme) D {
	pgv := counter.PgVals
	if len(s.radioBtns.Value) == 0 {
		s.radioBtns.Value = counter.INTEGERS
	}
	if s.radioBtns.Changed() {
		switch s.radioBtns.Value {
		case counter.PRIMES:
			pgv.SetActiveSequence(counter.PRIMES)
		case counter.FIBS:
			pgv.SetActiveSequence(counter.FIBS)
		case counter.NATURALS:
			pgv.SetActiveSequence(counter.NATURALS)
		case counter.INTEGERS:
			pgv.SetActiveSequence(counter.INTEGERS)
		}
		op.InvalidateOp{}.Add(gtx.Ops)
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
				layout.Rigid(material.RadioButton(th, &s.radioBtns, counter.INTEGERS, "Int").Layout),
				layout.Rigid(material.RadioButton(th, &s.radioBtns, counter.NATURALS, "Nat").Layout),
				layout.Rigid(material.RadioButton(th, &s.radioBtns, counter.PRIMES, "Primes").Layout),
				layout.Rigid(material.RadioButton(th, &s.radioBtns, counter.FIBS, "Fibs").Layout),
			)
		}),
	)
}
