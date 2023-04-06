package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var ASL1004 = &analysis.Analyzer{
	Name: "ASL1004",
	Doc:  "Verifies that log15 log string is capitalized.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		processCommonLog15Funcs(pass, false, func(fn string, e []ast.Expr) {
			if len(e) < 1 {
				return
			}

			switch argT := e[0].(type) {
			case *ast.BasicLit:
				if argT.Kind == token.STRING {
					argStr := removeStringQuotes(argT.Value)

					rarg := rune(argStr[0])
					if unicode.IsLower(rarg) && unicode.IsLetter(rarg) {
						report(pass, argT.Pos(), "ASL1004", "use of non-capitalized log string in log15 is not allowed")
					}
				}
			}
		})

		return nil, nil
	},
}

func processCommonLog15Funcs(pass *analysis.Pass, includeNew bool, processFn func(string, []ast.Expr)) {
	iteratePassFile(pass, func(f *ast.File) {
		ast.Inspect(f, func(n ast.Node) bool {
			switch nt := n.(type) {
			case *ast.CallExpr:
				switch fnt := nt.Fun.(type) {
				case *ast.SelectorExpr:
					switch fnt.Sel.Name {
					case "Crit":
					case "Error":
					case "Warn":
					case "Info":
					case "Debug":
					case "New":
						if !includeNew {
							return true
						}
					default:
						return true
					}

					isLog15Func := strings.HasSuffix(fmt.Sprint(pass.TypesInfo.Types[fnt.X].Type), "log15.Logger")

					if isLog15Func {
						processFn(fnt.Sel.Name, nt.Args)
					}
				}
			}

			return true
		})
	})
}
