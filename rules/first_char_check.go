package rules

import (
	"fmt"
	"go/ast"
	"strconv"
	"unicode"
)

type FirstCharChecker struct{}

func (FirstCharChecker) Name() string {
	return "first_char"
}

func (FirstCharChecker) Check(args []ast.Expr) error {
	var result = false

	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.BasicLit:
			unquoted, err := strconv.Unquote(v.Value)
			if err == nil {
				result = checkFirstLetter(unquoted)
			}
		}
		return false
	})

	if !result {
		return fmt.Errorf("log message should start with small letter")
	}
	return nil
}

func checkFirstLetter(s string) bool {
	if s == "" {
		return true
	}
	r := []rune(s)[0]
	return unicode.IsLower(r)
}
