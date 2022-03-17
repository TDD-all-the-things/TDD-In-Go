package stringcalculator

import "strconv"

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (s *StringCalculator) Add(numbers string) int {
	num, _ := strconv.Atoi(numbers)
	return num
}
