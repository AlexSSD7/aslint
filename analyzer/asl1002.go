package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var ASL1002 = &analysis.Analyzer{
	Name: "ASL1002",
	Doc:  "Verifies that errors passed to fmt.Errorf() and errors.Wrap() are not capitalized.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		processCommonErrorFuncs(pass, func(arg ast.Expr) {
			switch argT := arg.(type) {
			case *ast.BasicLit:
				if argT.Kind == token.STRING {
					rarg := rune(argT.Value[1])
					if unicode.IsUpper(rarg) && unicode.IsLetter(rarg) {
						report(pass, arg.Pos(), "ASL1002", "use of capitalized string as error is not allowed")
					}
				}
			}
		})

		return nil, nil
	},
}

func processCommonErrorFuncs(pass *analysis.Pass, processFn func(ast.Expr)) {
	iteratePassFile(pass, func(f *ast.File) {
		ast.Inspect(f, func(n ast.Node) bool {
			switch nt := n.(type) {
			case *ast.CallExpr:
				fn, ok := nt.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}

				switch fn.Sel.String() {
				case "Errorf":
					if len(nt.Args) > 0 {
						processFn(nt.Args[0])
					}
				case "Wrap", "Wrapf":
					if len(nt.Args) > 1 {
						if !isErrType(pass.TypesInfo.Types[nt.Args[0]].Type) {
							return true
						}

						processFn(nt.Args[1])
					}
				case "New":
					if strings.Contains(fmt.Sprint(fn.X), "errors") && len(nt.Args) == 1 {
						// Catches errors.New(...)
						processFn(nt.Args[0])
					}
				}
			}

			return true
		})
	})
}
