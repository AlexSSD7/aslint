package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL1010 = &analysis.Analyzer{
	Name: "ASL1010",
	Doc:  "Verifies that log15 error logs contain `error` field.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1010", "log15 error logs should contain `error` field (NOT `err`, this is important)")
		}

		processCommonLog15Funcs(pass, func(fn string, e []ast.Expr) {
			if fn != "Error" {
				// We look for error log level only
				return
			}

			if len(e) < 1 {
				return
			}

			foundErr := false

			for i := 1; i < len(e); i++ {
				if i&1 == 1 {
					// Do this for every key
					switch argT := e[i].(type) {
					case *ast.BasicLit:
						if argT.Kind == token.STRING {
							argStr := removeStringQuotes(argT.Value)
							if argStr == "error" {
								foundErr = true
							}
						}
					}
				}
			}

			if !foundErr {
				reportFunc(e[len(e)-1].End())
			}
		})

		return nil, nil
	},
}
