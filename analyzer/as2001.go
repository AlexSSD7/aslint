package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var ASL2001 = &analysis.Analyzer{
	Name: "ASL2001",
	Doc:  "Verifies that no strconv.Itoa and strconv.Atoi are used",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL2001", "use strconv.FormatInt/ParseInt (or better, strconv.FormatUint/ParseUint when working with uint's) instead of strconv.Itoa/Atoi")
		}

		for _, f := range pass.Files {
			ast.Inspect(f, func(n ast.Node) bool {
				switch nt := n.(type) {
				case *ast.CallExpr:
					fn, ok := nt.Fun.(*ast.SelectorExpr)
					if !ok {
						return true
					}

					if !strings.Contains(fmt.Sprint(fn.X), "strconv") {
						return true
					}

					switch fn.Sel.String() {
					case "Atoi":
						reportFunc(fn.Sel.Pos())
					case "Itoa":
						reportFunc(fn.Sel.Pos())
					}
				}

				return true
			})
		}
		return nil, nil
	},
}
