package rules

import (
	"fmt"
	"unicode"
)

type EngChecker struct{}

func (EngChecker) Name() string {
	return "eng"
}

func (EngChecker) Check(s string) error {
	for _, symb := range s {
		if symb > unicode.MaxASCII {
			return fmt.Errorf("log message should be in English")
		}
	}
	return nil
}
