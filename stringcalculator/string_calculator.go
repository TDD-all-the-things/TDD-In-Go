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

func (s *StringCalculator) Add(template string) int {
	sum := 0
	delimiter, numbers := ",", template
	if strings.HasPrefix(template, `//`) {
		delimiter, numbers = ";", template[5:]
	}
	for _, number := range strings.Split(strings.ReplaceAll(numbers, `\n`, delimiter), delimiter) {
		num, _ := strconv.Atoi(number)
		sum += num
	}
	return sum
}
