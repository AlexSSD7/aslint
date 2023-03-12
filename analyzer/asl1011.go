package analyzer

// * Actually, I think we can use `log15.Crit` for config errors
// var ASL1011 = &analysis.Analyzer{
// 	Name: "ASL1011",
// 	Doc:  "Suggests to not use `log15.Crit` as nothing should be critical.",
// 	Run: func(pass *analysis.Pass) (interface{}, error) {
// 		reportFunc := func(pos token.Pos) {
// 			report(pass, pos, "ASL1011", "the application should not have places where an error is *critical*, and for configuration, a simple error should be used")
// 		}

// 		for _, f := range pass.Files {
// 			ast.Inspect(f, func(n ast.Node) bool {
// 				switch nt := n.(type) {
// 				case *ast.CallExpr:
// 					switch fnt := nt.Fun.(type) {
// 					case *ast.SelectorExpr:
// 						switch fnt.Sel.Name {
// 						case "Crit":
// 						default:
// 							return true
// 						}

// 						isLog15Func := strings.HasSuffix(fmt.Sprint(pass.TypesInfo.Types[fnt.X].Type), "log15.Logger")

// 						if isLog15Func {
// 							reportFunc(fnt.Pos())
// 						}
// 					}
// 				}

// 				return true
// 			})
// 		}

// 		return nil, nil
// 	},
// }
