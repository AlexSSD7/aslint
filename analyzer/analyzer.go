package analyzer

import (
	"github.com/kyoh86/nolint"
	"golang.org/x/tools/go/analysis"
)

type CombinedAnalyzer struct {
	analyzers []*analysis.Analyzer
}

func NewCombinedAnalyzer() *CombinedAnalyzer {
	return &CombinedAnalyzer{
		analyzers: []*analysis.Analyzer{
			ASL1001,
			ASL1002,
			ASL1003,
			ASL1004,
			ASL1005,
			ASL1006,
			ASL1007,
			ASL1008,
			ASL1009,
			ASL1010,
			// ASL1011,
			ASL2001,
			ASL2002,
		},
	}
}

func (ca *CombinedAnalyzer) Analyzers() []*analysis.Analyzer {
	return nolint.WrapAll(ca.analyzers...)
}
