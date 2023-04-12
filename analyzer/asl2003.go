package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL2003 = &analysis.Analyzer{
	Name: "ASL2003",
	Doc:  "Verifies that no unkeyed initializer values are provided when creating a struct.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL2003", "unkeyed struct initializer values are not allowed")
		}

		iteratePassFile(pass, func(f *ast.File) {
			ast.Inspect(f, func(np ast.Node) bool {
				if np == nil {
					return true
				}

				switch n := np.(type) {
				case *ast.CompositeLit:
					switch n.Type.(type) {
					case *ast.Ident:
					default:
						return true
					}

					hasUnkeyed := false

					for _, elt := range n.Elts {
						if _, ok := elt.(*ast.KeyValueExpr); !ok {
							hasUnkeyed = true
							break
						}
					}

					if hasUnkeyed {
						reportFunc(n.Pos())
					}
				}

				return true
			})
		})

		return nil, nil
	},
}
