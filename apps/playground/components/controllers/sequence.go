package controllers

import (
	"gioui-experiment/apps/playground/components/data"
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
	cv := data.CurrVals
	if len(s.radioBtns.Value) == 0 {
		s.radioBtns.Value = data.INTEGERS
	}
	if s.radioBtns.Changed() {
		switch s.radioBtns.Value {
		case data.PRIMES:
			cv.SetActiveSequence(data.PRIMES)
		case data.FIBS:
			cv.SetActiveSequence(data.FIBS)
		case data.NATURALS:
			cv.SetActiveSequence(data.NATURALS)
		case data.INTEGERS:
			cv.SetActiveSequence(data.INTEGERS)
		}
		op.InvalidateOp{}.Add(gtx.Ops)
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
				layout.Rigid(material.RadioButton(th, &s.radioBtns, data.INTEGERS, "Int").Layout),
				layout.Rigid(material.RadioButton(th, &s.radioBtns, data.NATURALS, "Nat").Layout),
				layout.Rigid(material.RadioButton(th, &s.radioBtns, data.PRIMES, "Primes").Layout),
				layout.Rigid(material.RadioButton(th, &s.radioBtns, data.FIBS, "Fibs").Layout),
			)
		}),
	)
}
