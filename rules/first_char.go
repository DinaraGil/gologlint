package rules

import (
	"fmt"
)

type FirstCharChecker struct{}

func (FirstCharChecker) Name() string {
	return "first_char"
}

func (FirstCharChecker) Check(s string) error {
	if s == "" {
		return nil
	}
	if 'A' <= s[0] && s[0] <= 'Z' {
		return fmt.Errorf("log message should start with small letter")
	}
	//r := []rune(s)[0]
	//if unicode.IsUpper(r) {
	//	return fmt.Errorf("log message should start with small letter")
	//}
	return nil
}
