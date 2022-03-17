package stringcalculator

import (
	"regexp"
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
	reg, err := regexp.Compile(`^//[\D]+\\n`)
	if err != nil {
		return
	}
	b := []byte(template)
	loc := reg.FindIndex(b)
	if len(loc) == 0 {
		return
	}
	l, r := loc[0]+len(`//`), loc[1]-len(`\n`)
	return string(b[l:r]), string(b[loc[1]:])
}
