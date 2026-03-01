package rules

import (
	"go/ast"
	"strconv"
	"unicode"

	"github.com/DinaraGil/gologlint/core"
)

type SpecialSymbolsChecker struct{}

func (SpecialSymbolsChecker) Name() string {
	return "special symbols"
}

func specialSymbolsCheck(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r)) {
			return false
		}
	}
	return true
}

func (SpecialSymbolsChecker) Check(args []ast.Expr) *core.LintError {
	var result = false
	var noLit = true

	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.BasicLit:
			noLit = false
			unquoted, err := strconv.Unquote(v.Value)
			if err == nil {
				result = specialSymbolsCheck(unquoted)
			} else {
				result = specialSymbolsCheck(v.Value)
			}
			if !result {
				return false
			}
		}
		return true
	})

	if noLit || result {
		return nil
	}

	return &core.LintError{
		Rule: SpecialSymbolsChecker{}.Name(),
		Msg:  "log message should not contain special symbols",
	}
}
