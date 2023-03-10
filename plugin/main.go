package main

import (
	"github.com/AlexSSD7/aslint/analyzer"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return analyzer.NewCombinedAnalyzer().Analyzers()
}

var AnalyzerPlugin analyzerPlugin
