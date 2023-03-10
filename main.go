package main

import (
	"github.com/AlexSSD7/aslint/analyzer"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(analyzer.NewCombinedAnalyzer().Analyzers()...)
}
