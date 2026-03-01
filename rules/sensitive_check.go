package rules

import (
	"go/ast"
	"regexp"
	"strconv"
	"strings"

	"github.com/DinaraGil/gologlint/core"
)

type SensitiveChecker struct{}

func (SensitiveChecker) Name() string {
	return "sensitive check"
}

var sensitivePattern = regexp.MustCompile(`\b(password|token|secret|api[_]?key)\b`)

func sensitiveCheck(s string) bool {
	return sensitivePattern.MatchString(strings.ToLower(s))
}

func (SensitiveChecker) Check(args []ast.Expr) *core.LintError {
	var identHasSensitive = false
	var basicLitHasSensitive = false
	var hasIdent = false

	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.BasicLit:
			unquoted, err := strconv.Unquote(v.Value)
			if err != nil {
				basicLitHasSensitive = sensitiveCheck(v.Value)
			} else {
				basicLitHasSensitive = sensitiveCheck(unquoted)
			}
		case *ast.Ident:
			hasIdent = true
			identHasSensitive = sensitiveCheck(v.Name)
			if identHasSensitive {
				return false
			}
		}
		return true
	})

	if hasIdent && (basicLitHasSensitive || identHasSensitive) {
		return &core.LintError{
			Rule: SensitiveChecker{}.Name(),
			Msg:  "log message should not contain sensitive data",
		}
	}
	return nil
}
