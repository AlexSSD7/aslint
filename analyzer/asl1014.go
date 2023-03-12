package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL1014 = &analysis.Analyzer{
	Name: "ASL1014",
	Doc:  "Verifies that log15 logs contain at least one context field.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1014", "log15 log should contain at least one context field")
		}

		processCommonLog15Funcs(pass, func(fn string, e []ast.Expr) {
			if len(e) == 1 {
				reportFunc(e[0].Pos())
			}

		})

		return nil, nil
	},
}
