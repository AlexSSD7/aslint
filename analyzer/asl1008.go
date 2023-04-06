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

		processCommonLog15Funcs(pass, true, func(fn string, e []ast.Expr) {
			if len(e) < 2 {
				return
			}

			order := 1

			if fn == "New" {
				order = 0
			}

			for i := order; i < len(e); i++ {
				if i&1 == order {
					// Do this for every log key while omitting log values
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
