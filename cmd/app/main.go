package main

import (
	"fmt"
	"fuzzycontroller/internal/fuzzy"
)

func main() {

	FuzzyCntl := fuzzy.NewFuzzyController()
	concl := FuzzyCntl.GetFuzzyConclusion(fuzzy.F_LN, fuzzy.F_SP)
	fmt.Printf("%f", concl)
}
