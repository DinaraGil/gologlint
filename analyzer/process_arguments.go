package analyzer

import (
	"fmt"
	"github.com/DinaraGil/gologlint/rules"
	"go/ast"
)

//func convertToString(arg ast.Expr) string {
//	switch v := arg.(type) {
//	case *ast.BinaryExpr:
//		return convertToString(v.X) + convertToString(v.Y)
//	case *ast.BasicLit:
//		unquated, err := strconv.Unquote(v.Value)
//		if err == nil {
//			return unquated
//		}
//		return v.Value
//	case *ast.Ident:
//		return v.Obj.Name
//	default:
//		return fmt.Sprintf("чхз, %T", v)
//	}
//	return ""
//}

//
//func parseArgs(args []ast.Expr) string {
//	var strRes string
//	for _, arg := range args {
//		strRes += convertToString(arg)
//	}
//	return strRes
//}

func check(args []ast.Expr) []error {
	checkers := rules.GetCheckers()
	fmt.Printf("%s, %T", args, args)
	//stringArgs := parseArgs(args)

	return RunChecks(args, checkers)
}
