package analyzer

import (
	"go/ast"
	"strings"

	"go.uber.org/zap/zapcore"
	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		//inspect - dfs of nodes
		ast.Inspect(file, func(n ast.Node) bool {
			node, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			if selExpr, ok := node.Fun.(*ast.SelectorExpr); ok {
				obj := pass.TypesInfo.ObjectOf(selExpr.Sel)
				level, ok := zapcore.ParseLevel(strings.ToLower(obj.Name()))
				if ok != nil {
					return true
				}

				switch obj.Pkg().Name() {
				case "log", "slog", "zap":
					pass.Reportf(node.Pos(), "  TypesInfo.ObjectOf: %v (%T)", obj, obj)
					pass.Reportf(node.Pos(), "    Pkg: %v", obj.Pkg())
					pass.Reportf(node.Pos(), "    Name: %v", obj.Name())
					pass.Reportf(node.Pos(), "    Type: %v", obj.Type())
					pass.Reportf(node.Pos(), "parse level: %s", level)
					pass.Reportf(node.Pos(), "msg: %s", node.Args[0].(*ast.BasicLit).Value) //передавать ARGS
				default:
					return true
				}
			}
			return true
		})
	}
	return nil, nil
}
