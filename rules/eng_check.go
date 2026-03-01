package rules

import (
	"go/ast"
	"strconv"
	"unicode"

	"github.com/DinaraGil/gologlint/core"
)

type EngChecker struct{}

func (EngChecker) Name() string {
	return "eng"
}

func engCheck(s string) bool {
	for _, symb := range s {
		if symb > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func (EngChecker) Check(args []ast.Expr) *core.LintError {
	var result = false
	var noLit = true
	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.BasicLit:
			noLit = false
			unquoted, err := strconv.Unquote(v.Value)
			if err == nil {
				result = engCheck(unquoted)
			} else {
				result = engCheck(v.Value)
			}
			if !result {
				return false
			}
		}
		return true
	})

	if result || noLit {
		return nil
	}
	return &core.LintError{
		Rule: FirstCharChecker{}.Name(),
		Msg:  "log message should contain only English letters",
	}
}
