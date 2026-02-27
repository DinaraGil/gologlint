package rules

import "go/ast"

func treeIteration(arg ast.Expr, checkArg func(ast.Expr) bool) bool {
	switch v := arg.(type) {
	case *ast.BinaryExpr:
		return treeIteration(v.X, checkArg) && treeIteration(v.Y, checkArg)
	default:
		return checkArg(v)
	}
}

func IterateArgs(args []ast.Expr, checkArg func(ast.Expr) bool) {
	for _, arg := range args {
		if !treeIteration(arg, checkArg) {
			return
		}
	}
}
