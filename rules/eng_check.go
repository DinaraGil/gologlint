package rules

import (
	"fmt"
	"go/ast"
	"strconv"
	"unicode"
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

func (EngChecker) Check(args []ast.Expr) error {
	var result = false

	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.Ident:
			result = true
			return false
		case *ast.BasicLit:
			unquoted, err := strconv.Unquote(v.Value)
			if err == nil {
				result = engCheck(unquoted)
			}
		}
		return false
	})

	if !result {
		return fmt.Errorf("log message should contain only English letters")
	}
	return nil
}
