package analyzer

import "github.com/DinaraGil/gologlint/core"

func RunChecks(input string, checkers []core.Checker) []error {
	var errs []error

	for _, c := range checkers {
		if err := c.Check(input); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
