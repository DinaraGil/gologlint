package tests

import (
	"testing"

	"github.com/DinaraGil/gologlint/rules"
)

func TestFirstCharChecker(t *testing.T) {
	tests := []TestStruct{
		{
			name: "uppercase log message should return error",
			code: `package main
			func main() {
				log.Info("Starting server on port 8080")
			}`,
			expectingError: true,
		},
		{
			name: "uppercase log message should return error",
			code: `package main
			func main() {
				slog.Error("Failed to connect to database")
			}`,
			expectingError: true,
		},
		{
			name: "uppercase log message should not return error",
			code: `package main
			func main() {
				log.Info("starting server on port 8080")
			}`,
			expectingError: false,
		},
		{
			name: "uppercase log message should not return error",
			code: `package main
			func main() {
				slog.Error("failed to connect to database")
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

	checker := rules.FirstCharChecker{}

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
