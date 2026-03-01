package tests

import (
	"testing"

	"github.com/DinaraGil/gologlint/rules"
)

func TestSpecialCharChecker(t *testing.T) {
	tests := []TestStruct{
		{
			name: "emoji in log message should return error",
			code: `package main
			func main() {
				log.Info("server started!🚀")
			}`,
			expectingError: true,
		},
		{
			name: "special symbol in log message should return error",
			code: `package main
			func main() {
				log.Error("connection failed!!!")
			}`,
			expectingError: true,
		},
		{
			name: "special symbols in log message should not return error",
			code: `package main
			func main() {
				log.Warn("warning: something went wrong...")
			}`,
			expectingError: true,
		},
		{
			name: "only special characters should return error",
			code: `package main
				func main() {
					log.Info("!!!")
				}`,
			expectingError: true,
		},
		{
			name: "log message without special symbols should not return error",
			code: `package main
			func main() {
				log.Info("server started")
			}`,
			expectingError: false,
		},
		{
			name: "log message without special symbols should not return error",
			code: `package main
			func main() {
				log.Error("connection failed")
			}`,
			expectingError: false,
		},
		{
			name: "log message without special symbols should not return error",
			code: `package main
			func main() {
				log.Warn("something went wrong")
			}`,
			expectingError: false,
		},
		{
			name: "empty log message should not return error",
			code: `package main
				func main() {
					log.Info("")
				}`,
			expectingError: false,
		},
	}

	checker := rules.SpecialSymbolsChecker{}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			args := extractArgs(t, testCase.code)

			lintErr := checker.Check(args)

			if testCase.expectingError && lintErr == nil {
				t.Errorf("expected error, got nil")
			}
			if !testCase.expectingError && lintErr != nil {
				t.Errorf("unexpected error: %s", lintErr.Msg)
			}
		})
	}
}
