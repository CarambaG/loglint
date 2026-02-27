package analyzer

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func checkLowercase(msg string) bool {
	if msg == "" {
		return true
	}

	r := []rune(msg)[0]
	return unicode.IsLower(r)
}

func isEnglish(msg string) bool {
	for _, r := range msg {
		if unicode.In(r, unicode.Cyrillic) {
			return false
		}
	}
	return true
}

func hasSpecialChars(msg string) bool {
	for _, r := range msg {
		if unicode.IsPunct(r) && r != '.' {
			return true
		}

		if r > 0x1F600 { // проверка emoji
			return true
		}
	}
	return false
}

var sensitiveKeywords = []string{
	"password",
	"token",
	"api_key",
	"secret",
}

func containsSensitive(msg string) bool {
	l := strings.ToLower(msg)

	for _, kw := range sensitiveKeywords {
		if strings.Contains(l, kw) {
			return true
		}
	}
	return false
}

func checkRules(pass *analysis.Pass, call *ast.CallExpr, msg string) {

	if !checkLowercase(msg) {
		pass.Reportf(call.Pos(), "log message must start with lowercase letter")
	}

	if !isEnglish(msg) {
		pass.Reportf(call.Pos(), "log message must be in English")
	}

	if hasSpecialChars(msg) {
		pass.Reportf(call.Pos(), "log message must not contain special characters or emoji")
	}

	if containsSensitive(msg) {
		pass.Reportf(call.Pos(), "log message contains potentially sensitive data")
	}
}
