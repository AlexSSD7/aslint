package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var ASL1003 = &analysis.Analyzer{
	Name: "ASL1003",
	Doc:  "Verifies that errors passed to fmt.Errorf() and errors.Wrap() do not contain common 'failed', 'unable', 'cannot' keywords that make chained error handling a nightmare.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		processCommonErrorFuncs(pass, func(arg ast.Expr) {
			switch argT := arg.(type) {
			case *ast.BasicLit:
				if argT.Kind == token.STRING {
					argStr := removeStringQuotes(argT.Value)

					contains := false

					contains = contains || strings.HasPrefix(argStr, "failed")
					contains = contains || strings.HasPrefix(argStr, "unable")
					contains = contains || strings.HasPrefix(argStr, "can't")
					contains = contains || strings.HasPrefix(argStr, "cant")
					contains = contains || strings.HasPrefix(argStr, "cannot")
					contains = contains || strings.HasPrefix(argStr, "could not")

					if contains {
						report(pass, arg.Pos(), "ASL1003", "use of generic error keywords like 'failed', 'unable', 'cannot' and similar is not allowed")
					}
				}
			}
		})

		return nil, nil
	},
}
