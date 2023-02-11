package meander

import (
	"errors"
	"strings"
)

func ParseCost(s string) Cost {
	return costStrings[s]
}

type CostRange struct {
	From Cost
	To   Cost
}

func ParseCostRange(s string) (CostRange, error) {
	var result CostRange
	sArr := strings.Split(s, "...")
	if len(sArr) != 2 {
		return result, errors.New("invalid cost range")
	}

	result.From = ParseCost(sArr[0])
	result.To = ParseCost(sArr[1])
	return result, nil
}
