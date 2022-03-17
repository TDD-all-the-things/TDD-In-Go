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
	sum := 0
	for _, number := range strings.Split(numbers, ",") {
		num, _ := strconv.Atoi(number)
		sum += num
	}
	return sum
}
