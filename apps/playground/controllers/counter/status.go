package counter

import (
	"gioui-experiment/apps/playground/data"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"strconv"
)

type Status struct {
	data.Generator
	primesState   component.DiscloserState
	fibsState     component.DiscloserState
	naturalsState component.DiscloserState
	integersState component.DiscloserState
}

// Layout - TODO: hardcoded due to UTTER LAZINESS
// More "dynamics" to be added
func (s *Status) Layout(gtx C, th *material.Theme) D {
	pgv := data.PgVals

	// TODO: rethink location of Cache population
	if len(pgv.Cache[data.PRIMES]) != data.PLIMIT {
		s.GenPrimes(data.PLIMIT)
	}
	if len(pgv.Cache[data.FIBS]) != data.FLIMIT {
		s.GenFibs(data.FLIMIT)
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Inset{
				Bottom: unit.Dp(5),
			}.Layout(gtx, func(gtx C) D {
				return component.SimpleDiscloser(th, &s.integersState).Layout(gtx,
					material.Body1(th, "Integers").Layout,
					func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Displayed").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Integers.Displayed, 10)).Layout(gtx)
									}))
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Start").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Integers.Start, 10)).Layout(gtx)
									}))
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Step").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Integers.Step, 10)).Layout(gtx)
									}))
							}),
						)
					},
				)
			})
		}),

		layout.Rigid(func(gtx C) D {
			return layout.Inset{
				Top:    unit.Dp(5),
				Bottom: unit.Dp(5),
			}.Layout(gtx, func(gtx C) D {
				return component.SimpleDiscloser(th, &s.naturalsState).Layout(gtx,
					material.Body1(th, "Naturals").Layout,
					func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Displayed").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Naturals.Displayed, 10)).Layout(gtx)
									}))
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Start").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Naturals.Start, 10)).Layout(gtx)
									}))
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Step").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Naturals.Step, 10)).Layout(gtx)
									}))
							}),
						)
					},
				)
			})
		}),

		layout.Rigid(func(gtx C) D {
			return layout.Inset{
				Top:    unit.Dp(5),
				Bottom: unit.Dp(5),
			}.Layout(gtx, func(gtx C) D {
				return component.SimpleDiscloser(th, &s.primesState).Layout(gtx,
					material.Body1(th, "Primes").Layout,
					func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Cached").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(uint64(len(pgv.Cache["primes"])), 10)).Layout(gtx)
									}))
							}),

							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Position").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatInt(int64(pgv.Primes.Index+1), 10)).Layout(gtx)
									}))
							}),

							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Displayed").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Cache["primes"][pgv.Primes.Index], 10)).Layout(gtx)
									}))
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Start").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Primes.Start, 10)).Layout(gtx)
									}))
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Step").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Primes.Step, 10)).Layout(gtx)
									}))
							}),
						)
					})
			})
		}),

		layout.Rigid(func(gtx C) D {
			return layout.Inset{
				Top: unit.Dp(5),
			}.Layout(gtx, func(gtx C) D {
				return component.SimpleDiscloser(th, &s.fibsState).Layout(gtx,
					material.Body1(th, "Fibonacci").Layout,
					func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Cached").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(uint64(len(pgv.Cache["fibs"])), 10)).Layout(gtx)
									}))
							}),

							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Position").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatInt(int64(pgv.Fibonacci.Index+1), 10)).Layout(gtx)
									}))
							}),

							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Displayed").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Cache["fibs"][pgv.Fibonacci.Index], 10)).Layout(gtx)
									}))
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Start").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Fibonacci.Start, 10)).Layout(gtx)
									}))
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, "Step").Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, strconv.FormatUint(pgv.Fibonacci.Step, 10)).Layout(gtx)
									}))
							}),
						)
					})
			})
		}))
}
