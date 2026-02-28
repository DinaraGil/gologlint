package analyzer

import (
	"go/ast"

	"github.com/DinaraGil/gologlint/rules"
	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	var checkers = rules.GetCheckers()

	for _, file := range pass.Files {
		//inspect - dfs of nodes
		ast.Inspect(file, func(n ast.Node) bool {
			node, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			selExpr, ok := node.Fun.(*ast.SelectorExpr)

			if !ok {
				return true
			}

			obj := pass.TypesInfo.ObjectOf(selExpr.Sel)

			if !IsLoggingCall(obj) {
				return true
			}

			lintErrors := checkLogArgs(checkers, node.Args)

			for _, lintErr := range lintErrors {
				pass.Reportf(node.Pos(), "%s: %s", lintErr.Rule, lintErr.Msg)
			}
			return true
		})
	}
	return nil, nil
}
