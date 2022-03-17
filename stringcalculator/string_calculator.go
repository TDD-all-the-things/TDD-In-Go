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
		sum := 0
		for _, number := range strings.SplitN(numbers, ",", 2) {
			num, _ := strconv.Atoi(number)
			sum += num
		}
		return sum
	}
	num, _ := strconv.Atoi(numbers)
	return num
}
