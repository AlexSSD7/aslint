package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var ASL1013 = &analysis.Analyzer{
	Name: "ASL1013",
	Doc:  "Verifies that log15 logs contain no complex punctuation.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1013", "log15 should not contain complex punctuation")
		}

		processCommonLog15Funcs(pass, func(fn string, e []ast.Expr) {
			if len(e) < 1 {
				return
			}

			switch argT := e[0].(type) {
			case *ast.BasicLit:
				if argT.Kind == token.STRING {
					argStr := strings.ToLower(removeStringQuotes(argT.Value))

					contains := false

					contains = contains || strings.Contains(argStr, "\"")
					contains = contains || strings.Contains(argStr, "'")
					contains = contains || strings.Contains(argStr, "`")
					contains = contains || strings.Contains(argStr, "--")
					contains = contains || strings.Contains(argStr, ".")
					contains = contains || strings.Contains(argStr, ";")
					contains = contains || strings.Contains(argStr, ":")
					contains = contains || strings.Contains(argStr, "!")

					if contains {
						reportFunc(argT.Pos())
					}
				}
			}
		})

		return nil, nil
	},
}
