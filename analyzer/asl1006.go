package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL1006 = &analysis.Analyzer{
	Name: "ASL1006",
	Doc:  "Verifies that no strings are concatenated in fmt.Errorf().",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		reportFunc := func(pos token.Pos) {
			report(pass, pos, "ASL1006", "strings should not be concatenated in fmt.Errorf(), please use formatting directives (like '%v') and supply values as parameters instead")
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
					case "Errorf":
						if len(nt.Args) > 0 {
							processFn(nt.Args[0])
						}
					case "Wrapf":
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
