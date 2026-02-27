package rules

import (
	"github.com/DinaraGil/gologlint/core"
)

func GetCheckers() []core.Checker {
	return []core.Checker{
		FirstCharChecker{},
		EngChecker{},
		//SpecialSymbolsChecker{},
		//SensitiveChecker{},
	}
}
