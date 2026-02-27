package rules

import (
	"fmt"
	"unicode"
)

type SpecialSymbolsChecker struct{}

func (SpecialSymbolsChecker) Name() string {
	return "special symbols"
}

func (SpecialSymbolsChecker) Check(s string) error {
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r)) {
			return fmt.Errorf("log message should not contain special symbols")
		}
	}
	return nil
}
