package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages format",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {

			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isSupportedLogger(pass, call) {
				return true
			}

			msg := extractLogMessage(call)
			if msg == "" {
				return true
			}

			checkRules(pass, call, msg)

			return true
		})
	}

	return nil, nil
}

func extractLogMessage(call *ast.CallExpr) string {

	if len(call.Args) == 0 {
		return ""
	}

	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok {
		return ""
	}

	if lit.Kind != token.STRING {
		return ""
	}

	msg := strings.Trim(lit.Value, `"`)

	return msg
}

func isLogCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.Uses[selector.Sel]
	if obj == nil {
		return false
	}

	if obj.Pkg() == nil {
		return false
	}

	return obj.Pkg().Path() == "log"
}

func isSlogCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.Uses[selector.Sel]
	if obj == nil {
		return false
	}

	if obj.Pkg() == nil {
		return false
	}

	return obj.Pkg().Path() == "log/slog"
}

func isZapCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	typ := pass.TypesInfo.TypeOf(selector.X)
	if typ == nil {
		return false
	}

	return strings.Contains(typ.String(), "zap.Logger")
}

func isSupportedLogger(pass *analysis.Pass, call *ast.CallExpr) bool {
	return isSlogCall(pass, call) || isZapCall(pass, call) || isLogCall(pass, call)
}
