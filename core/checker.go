package core

import "go/ast"

type Checker interface {
	Check([]ast.Expr) error
	Name() string
}
