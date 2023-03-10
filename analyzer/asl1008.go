package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var ASL1008 = &analysis.Analyzer{
	Name: "ASL1008",
	Doc:  "Verifies that log15 context field keys are using dash-case.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1008", "use of any other case (like camelCase or snake_case) other than dash-case in log15 context keys is not allowed")
		}

		processCommonLog15Funcs(pass, func(fn string, e []ast.Expr) {
			if len(e) < 2 {
				return
			}

			for i := 1; i < len(e); i++ {
				if i&1 == 1 {
					// Do this for every key
					switch argT := e[i].(type) {
					case *ast.BasicLit:
						if argT.Kind == token.STRING {
							argStr := removeStringQuotes(argT.Value)
							if strings.ToLower(argStr) != argStr || strings.Contains(argStr, "_") {
								reportFunc(argT.Pos())
							}
						}
					}
				}
			}
		})

		return nil, nil
	},
}
