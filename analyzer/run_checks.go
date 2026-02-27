package analyzer

import (
	"go/ast"

	"github.com/DinaraGil/gologlint/core"
)

func RunChecks(args []ast.Expr, checkers []core.Checker) []error {
	var errs []error

	for _, c := range checkers {
		if err := c.Check(args); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
