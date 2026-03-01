package tests

import (
	"testing"

	"github.com/DinaraGil/gologlint/rules"
)

func TestSensitiveCharChecker(t *testing.T) {
	tests := []TestStruct{
		{
			name: "password word in log message should return error",
			code: `package main
			func main() {
				log.Info("user password: " + password)
			}`,
			expectingError: true,
		},
		{
			name: "word apikey in log message should return error",
			code: `package main
			func main() {
				log.Debug("api_key=" + apiKey)
			}`,
			expectingError: true,
		},
		{
			name: "word token in log message should not return error",
			code: `package main
			func main() {
				log.Info("token: " + token)
			}`,
			expectingError: true,
		},
		{
			name: "log message without ident should not return error",
			code: `package main
			func main() {
				log.Info("user authenticated successfully")
			}`,
			expectingError: false,
		},
		{
			name: "log message without ident should not return error",
			code: `package main
			func main() {
				log.Debug("api request completed")
			}`,
			expectingError: false,
		},
		{
			name: "log message without ident should not return error",
			code: `package main
			func main() {
				log.Info("token validated")
			}`,
			expectingError: false,
		},
	}

	checker := rules.SensitiveChecker{}

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
