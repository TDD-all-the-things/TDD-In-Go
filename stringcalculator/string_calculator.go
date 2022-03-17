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
	delimiter, numbers := s.parseTemplate(template)
	for _, number := range strings.Split(strings.ReplaceAll(numbers, `\n`, delimiter), delimiter) {
		num, _ := strconv.Atoi(number)
		sum += num
	}
	return sum
}

func (s *StringCalculator) parseTemplate(template string) (delimiter string, numbers string) {
	delimiter, numbers = ",", template
	if strings.HasPrefix(template, `//`) {
		delimiter, numbers = ";", template[5:]
	}
	return
}
