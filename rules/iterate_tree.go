package rules

import "go/ast"

func iterateArg(arg ast.Expr, processArg func(ast.Expr) bool) bool {
	switch v := arg.(type) {
	case *ast.BinaryExpr:
		return iterateArg(v.X, processArg) && iterateArg(v.Y, processArg)
	default:
		return processArg(v)
	}
}

func IterateArgs(args []ast.Expr, processArg func(ast.Expr) bool) {
	for _, arg := range args {
		if !iterateArg(arg, processArg) {
			return
		}
	}
}
