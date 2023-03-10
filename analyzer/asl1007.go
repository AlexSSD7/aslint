package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL1007 = &analysis.Analyzer{
	Name: "ASL1007",
	Doc:  "Verifies that the error message is not concatenated in errors.Wrap().",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1007", "error message should not be concatenated in errors.Wrap(), use errors.Wrapf() with formatting directives (like '%v') instead")
		}

		processFn := func(e ast.Expr) {
			switch argT := e.(type) {
			case *ast.BinaryExpr:
				reportFunc(argT.Pos())
			}
		}

		iteratePassFile(pass, func(f *ast.File) {
			ast.Inspect(f, func(n ast.Node) bool {
				switch nt := n.(type) {
				case *ast.CallExpr:
					fn, ok := nt.Fun.(*ast.SelectorExpr)
					if !ok {
						return true
					}

					switch fn.Sel.String() {
					case "Wrap":
						if len(nt.Args) > 1 {
							if !isErrType(pass.TypesInfo.Types[nt.Args[0]].Type) {
								return true
							}

							processFn(nt.Args[1])
						}
					}
				}

				return true
			})
		})

		return nil, nil
	},
}
