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
	replacedNumbers := strings.ReplaceAll(numbers, `\n`, ",")
	for _, number := range strings.Split(replacedNumbers, ",") {
		num, _ := strconv.Atoi(number)
		sum += num
	}
	return sum
}
