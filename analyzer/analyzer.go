package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "gologlint",
	Doc:  "check if the log is ok",
	Run:  run,
}
