package fuzzy

// linguistic vars with there fuzzy values
const (
	F_VLN = -1.0  // Very Large Negative
	F_LN  = -0.72 // Large Negative
	F_MN  = -0.44 // Meduim Negative
	F_SN  = -0.16 // Small Negative
	F_NO  = 0.0
	F_SP  = 0.16
	F_MP  = 0.44
	F_LP  = 0.72
	F_VLP = 1.0
)

type FuzzyController struct {
	fcFuzzy Fuzzy
}

// Fuzzy regulator solution. Fuzzy Conclusion
// Returning impact
//
// e current error,
// de first derivative of errors
func (fc *FuzzyController) GetFuzzyConclusion(e float64, de float64) float64 {
	summ_alpha_c := 0.0
	summ_alpha := 0.0

	// Composite all Fuzzy Rules and Centriod Defuzzification
	for _, rule := range fc.fcFuzzy.Rules {

		alpha := 0.0
		mue := 0.0
		mude := 0.0

		mue = fc.fcFuzzy.Gmu(e, rule.ferr)
		mude = fc.fcFuzzy.Gmu(de, rule.fderr)

		alpha = rule.op.ft(mue, mude)
		summ_alpha_c += (alpha * rule.fthen)
		summ_alpha += alpha
	}

	return summ_alpha_c / summ_alpha
}

func NewFuzzyController() *FuzzyController {

	lfcFuzzy := Fuzzy{}

	// all is ok
	lfcFuzzy.addRule(F_NO, new(F_and), F_NO, F_NO)

	// the last impact does not bring results
	lfcFuzzy.addRule(F_VLN, new(F_or), F_VLN, F_VLP)
	lfcFuzzy.addRule(F_VLP, new(F_or), F_VLP, F_VLN)

	// we have positive effects from last impact
	lfcFuzzy.addRule(F_SN, new(F_and), F_SN, F_SP)
	lfcFuzzy.addRule(F_SP, new(F_and), F_SP, F_SN)
	lfcFuzzy.addRule(F_LN, new(F_and), F_SN, F_LP)
	lfcFuzzy.addRule(F_LP, new(F_and), F_SP, F_LN)

	// situation had become worse
	lfcFuzzy.addRule(F_LN, new(F_and), F_SP, F_LP)
	lfcFuzzy.addRule(F_LP, new(F_and), F_SN, F_LN)

	// and so on
	// ...

	return &FuzzyController{fcFuzzy: lfcFuzzy}
}
