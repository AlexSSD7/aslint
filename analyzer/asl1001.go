package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var ASL1001 = &analysis.Analyzer{
	Name: "ASL1001",
	Doc:  "Verifies that no plain errors (without message and/or chain) are returned.",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		processFn := func(pass *analysis.Pass, block ast.Node) {
			var noProcessFuncsPositions []token.Pos

			ast.Inspect(block, func(n ast.Node) bool {
				switch item := n.(type) {
				case *ast.ReturnStmt:
					for _, retResult := range item.Results {
						if !isErrType(pass.TypesInfo.Types[retResult].Type) {
							continue
						}

						_, usesProcessFunc := retResult.(*ast.CallExpr)

						if !usesProcessFunc {
							noProcessFuncsPositions = append(noProcessFuncsPositions, retResult.Pos())
						}
					}
				}

				return true
			})

			if len(noProcessFuncsPositions) > 1 {
				for _, pos := range noProcessFuncsPositions {
					report(pass, pos, "ASL1001", "error chain/message missing, use errors.Wrap() or fmt.Errorf()")
				}
			}
		}

		iteratePassFile(pass, func(f *ast.File) {
			ast.Inspect(f, func(n ast.Node) bool {
				switch nt := n.(type) {
				case *ast.FuncDecl:
					processFn(pass, nt)
				case *ast.FuncLit:
					processFn(pass, nt)
				}

				return true
			})
		})

		return nil, nil
	},
}
