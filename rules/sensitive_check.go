package rules

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

type SensitiveChecker struct{}

func (SensitiveChecker) Name() string {
	return "sensitive check"
}

var sensitivePattern = regexp.MustCompile(`\b(password|token|secret|api[_]?key)\b`)

func sensitiveCheck(s string) bool {
	return sensitivePattern.MatchString(strings.ToLower(s))
}

func (SensitiveChecker) Check(args []ast.Expr) error {
	var result = false

	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.Ident:
			result = sensitiveCheck(v.Name)
			if result {
				return false
			}
		}
		return true
	})

	if result {
		return fmt.Errorf("log message should not contain sensitive data")
	}
	return nil
}
