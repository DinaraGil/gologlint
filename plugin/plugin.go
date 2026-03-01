package plugin

import (
	"github.com/DinaraGil/gologlint/analyzer"
	"golang.org/x/tools/go/analysis"
)

var Analyzers = []*analysis.Analyzer{
	analyzer.Analyzer,
}
