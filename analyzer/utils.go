package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func iteratePassFile(pass *analysis.Pass, fn func(*ast.File)) {
	for _, f := range pass.Files {
		// if shouldExcludeFile(f) {
		// 	continue
		// }
		f.Pos()

		fn(f)
	}
}

func report(pass *analysis.Pass, pos token.Pos, category string, message string) {
	pass.Report(analysis.Diagnostic{
		Pos:      pos,
		Category: category,
		Message:  message + ". If this is intentional, ignore with `//nolint:" + category + "`.",
	})
}

func isErrType(t types.Type) bool {
	return types.Implements(t, types.Universe.Lookup("error").Type().Underlying().(*types.Interface))
}

func removeStringQuotes(s string) string {
	if strings.HasPrefix(s, "\"") {
		return strings.TrimSuffix(s[1:], "\"")
	}

	if strings.HasPrefix(s, "`") {
		return strings.TrimSuffix(s[1:], "`")
	}

	return s
}
