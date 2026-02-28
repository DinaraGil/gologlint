package core

import "go/ast"

type Checker interface {
	Check([]ast.Expr) *LintError
	Name() string
}
