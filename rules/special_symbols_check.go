package rules

import (
	"fmt"
	"go/ast"
	"strconv"
	"unicode"
)

type SpecialSymbolsChecker struct{}

func (SpecialSymbolsChecker) Name() string {
	return "special symbols"
}

func speacialSymbolsCheck(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r)) {
			return false
		}
	}
	return true
}

func (SpecialSymbolsChecker) Check(args []ast.Expr) error {
	var result = false

	IterateArgs(args, func(arg ast.Expr) bool {
		switch v := arg.(type) {
		case *ast.BasicLit:
			unquoted, err := strconv.Unquote(v.Value)
			if err == nil {
				result = speacialSymbolsCheck(unquoted)
			}
		case *ast.Ident:
			result = true
			return false
		}
		return false
	})

	if !result {
		return fmt.Errorf("log message should not contain special symbols")
	}
	return nil
}
