package rules

import (
	"fmt"
	"go/ast"
	"unicode"
)

type FirstCharChecker struct{}

func (FirstCharChecker) Name() string {
	return "first_char"
}

func (FirstCharChecker) Check(args []ast.Expr) error {

	if s == "" {
		return nil
	}

	r := []rune(s)[0]
	if unicode.IsUpper(r) {
		return fmt.Errorf("log message should start with small letter")
	}
	return nil
}
