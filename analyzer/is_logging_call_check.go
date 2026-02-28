package analyzer

import (
	"go/types"
)

var validCalls = map[string]struct{}{
	"Debug": {},
	"Info":  {},
	"Warn":  {},
	"Error": {},
	"Fatal": {},
	"Panic": {},
}

var validPkgs = map[string]struct{}{
	"log":  {},
	"slog": {},
	"zap":  {},
}

func isValidPkg(pkgName string) bool {
	_, ok := validPkgs[pkgName]
	return ok
}

func isValidLevel(funcName string) bool {
	_, ok := validCalls[funcName]
	return ok
}

func IsLoggingCall(obj types.Object) bool {
	return isValidLevel(obj.Name()) && isValidPkg(obj.Pkg().Name())
}
