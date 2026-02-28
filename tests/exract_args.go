package tests

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func extractArgs(t *testing.T, src string) []ast.Expr {
	// FileSet хранит информацию о позициях в файле
	fset := token.NewFileSet()
	// Возвращает *astFile.File, mode=0 парсинг без дополнительных флагов
	astFile, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		t.Fatal(err)
	}

	var args []ast.Expr

	ast.Inspect(astFile, func(n ast.Node) bool {

		if call, ok := n.(*ast.CallExpr); ok {
			args = call.Args
			return false
		}
		return true
	})

	return args
}
