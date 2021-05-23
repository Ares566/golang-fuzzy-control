package fuzzy

import "math"

type Fuzzy struct {
	Rules []Rule
}

// Adding rule to controller
//
// ferr fuzzy value or error,
// op fuzzy logic operator,
// fderr fuzzy value of delta error (fderr = ferr - previous ferr),
// fthen resulting fuzzy value
func (f *Fuzzy) addRule(ferr float64, op Operator, fderr float64, fthen float64) {

	newRule := NewRule(ferr, op, fderr, fthen)
	f.Rules = append(f.Rules, *newRule)
}

// Gaussian Membership Function µ𝔸(x)
// The membership degree µ𝔸(x) quantifies the grade of membership of the element x to the fuzzy set.
//
// х element from Х
// A from fuzzy set 𝔸
// 0.3 basis
func (f *Fuzzy) Gmu(x float64, A float64) float64 {
	return math.Exp(-(math.Pow(x-A, 2) / (2 * math.Pow(0.3, 2))))
}
