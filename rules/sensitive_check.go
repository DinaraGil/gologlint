package rules

import (
	"fmt"
	"regexp"
	"strings"
)

type SensitiveChecker struct{}

func (SensitiveChecker) Name() string {
	return "sensitive check"
}

var sensitivePattern = regexp.MustCompile(`\b(password|token|secret|key)\b`)

func (SensitiveChecker) Check(s string) error {
	if sensitivePattern.MatchString(strings.ToLower(s)) {
		return fmt.Errorf("log message should not contain sensitive data")
	}

	return nil
}
