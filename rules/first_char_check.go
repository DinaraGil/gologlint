package rules

import (
	"go/ast"
	"strconv"
	"unicode"

	"github.com/DinaraGil/gologlint/core"
)

type FirstCharChecker struct{}

func (FirstCharChecker) Name() string {
	return "first_char"
}

func (FirstCharChecker) Check(args []ast.Expr) *core.LintError {
	var result = false
	var hasLit = false

	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.Ident:
			return true
		case *ast.BasicLit:
			hasLit = true
			unquoted, err := strconv.Unquote(v.Value)
			if err != nil {
				return true
			}
			result = checkFirstLetter(unquoted)
			if !result {
				return false
			}
		}
		return true
	})

	if result || !hasLit {
		return nil
	}

	return &core.LintError{
		Rule: FirstCharChecker{}.Name(),
		Msg:  "log message should start with small letter",
	}
}

func checkFirstLetter(s string) bool {
	if s == "" {
		return true
	}
	r := []rune(s)[0]
	return unicode.IsLower(r)
}
