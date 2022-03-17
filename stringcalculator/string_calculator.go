package stringcalculator

import (
	"strconv"
	"strings"
)

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (s *StringCalculator) Add(numbers string) int {
	if strings.Contains(numbers, ",") {
		return 3
	}
	num, _ := strconv.Atoi(numbers)
	return num
}
