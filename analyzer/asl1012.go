package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var ASL1012 = &analysis.Analyzer{
	Name: "ASL1012",
	Doc:  "Verifies that log15 error logs start with 'Failed' instead of 'Unable', 'Cannot', and similar.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1012", "log15 error logs should start with 'Failed', not 'Unable', 'Cannot', and similar")
		}

		processCommonLog15Funcs(pass, false, func(fn string, e []ast.Expr) {
			if fn != "Error" && fn != "Crit" {
				// We look for Error and Crit log levels only
				return
			}

			if len(e) < 1 {
				return
			}

			switch argT := e[0].(type) {
			case *ast.BasicLit:
				if argT.Kind == token.STRING {
					argStr := strings.ToLower(removeStringQuotes(argT.Value))

					contains := false

					contains = contains || strings.HasPrefix(argStr, "unable")
					contains = contains || strings.HasPrefix(argStr, "can't")
					contains = contains || strings.HasPrefix(argStr, "cant")
					contains = contains || strings.HasPrefix(argStr, "cannot")
					contains = contains || strings.HasPrefix(argStr, "could not")
					contains = contains || strings.HasPrefix(argStr, "couldn't")
					contains = contains || strings.HasPrefix(argStr, "error")
					contains = contains || strings.HasPrefix(argStr, "something went wrong")

					if contains {
						reportFunc(argT.Pos())
					}
				}
			}
		})

		return nil, nil
	},
}
