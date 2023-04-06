package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL1005 = &analysis.Analyzer{
	Name: "ASL1005",
	Doc:  "Verifies that log15 log string is static and is not modified.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1005", "use of non-static log string in log15 is not allowed, please use log fields")
		}
		processCommonLog15Funcs(pass, false, func(fn string, e []ast.Expr) {
			if len(e) < 1 {
				return
			}

			switch argT := e[0].(type) {
			case *ast.BasicLit:
				if argT.Kind != token.STRING {
					reportFunc(argT.Pos())
				}
			default:
				reportFunc(argT.Pos())
			}
		})
		return nil, nil
	},
}
