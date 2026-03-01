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
	var noLit = true

	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.Ident:
			return false
		case *ast.BasicLit:
			noLit = false
			unquoted, err := strconv.Unquote(v.Value)
			if err != nil {
				result = checkFirstLetter(v.Value)
			} else {
				result = checkFirstLetter(unquoted)
			}
			return false
		}
		return true
	})

	if result || noLit {
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
