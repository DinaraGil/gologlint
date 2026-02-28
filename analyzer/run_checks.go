package analyzer

import (
	"go/ast"

	"github.com/DinaraGil/gologlint/core"
)

func checkLogArgs(checkers []core.Checker, args []ast.Expr) (lintErrors []core.LintError) {
	for _, c := range checkers {
		checkRes := c.Check(args)
		if checkRes != nil {
			lintErrors = append(lintErrors, *checkRes)
		}
	}
	return
}
