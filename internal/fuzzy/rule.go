package fuzzy

type Rule struct {
	ferr  float64
	op    Operator
	fderr float64
	fthen float64
}

func NewRule(_ferr float64, _op Operator, _fderr float64, _fthen float64) *Rule {
	return &Rule{
		ferr:  _ferr,
		op:    _op,
		fderr: _fderr,
		fthen: _fthen,
	}
}
