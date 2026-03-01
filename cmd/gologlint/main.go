package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/DinaraGil/gologlint/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
