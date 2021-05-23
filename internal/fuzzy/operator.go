package fuzzy

import "math"

type Operator interface {
	ft(op1, op2 float64) float64
}

type F_and struct{}

func (f *F_and) ft(op1, op2 float64) float64 {
	return math.Min(op1, op2)
}

type F_or struct{}

func (f *F_or) ft(op1, op2 float64) float64 {
	return math.Max(op1, op2)
}
