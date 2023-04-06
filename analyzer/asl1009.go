package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL1009 = &analysis.Analyzer{
	Name: "ASL1009",
	Doc:  "Verifies that log15 context values are in non-odd length (as they are key-value pairs).",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1009", "the amount of log15 context values should be non-odd as context items are key-value pairs")
		}

		processCommonLog15Funcs(pass, false, func(fn string, e []ast.Expr) {
			if len(e) < 2 {
				return
			}

			if len(e)&1 != 1 {
				reportFunc(e[len(e)-1].End())
			}
		})

		return nil, nil
	},
}
