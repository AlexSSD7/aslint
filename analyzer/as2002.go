package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL2002 = &analysis.Analyzer{
	Name: "ASL2002",
	Doc:  "Verifies that no initializers are used declaring a variable with `var`.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL2002", "`var` should be used for empty (null) variable initializations only; in all other cases, please use `:=`")
		}

		iteratePassFile(pass, func(f *ast.File) {
			ast.Inspect(f, func(np ast.Node) bool {
				if np == nil {
					return true
				}

				switch np.(type) {
				case *ast.File:
					return true
				case *ast.FuncDecl, *ast.BlockStmt:
					ast.Inspect(np, func(n ast.Node) bool {
						if n == nil {
							return true
						}

						switch nt := n.(type) {
						case *ast.GenDecl:
							if nt.Tok == token.VAR {
								for _, spec := range nt.Specs {
									switch vs := spec.(type) {
									case *ast.ValueSpec:
										if len(vs.Values) != 0 {
											reportFunc(vs.Pos())
										}
									}
								}
							}
						}

						return true
					})
				}

				return false
			})
		})

		return nil, nil
	},
}
