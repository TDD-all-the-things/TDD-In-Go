package stringcalculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
)

type StringCalculator struct {
	addCalledCounter int32
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (s *StringCalculator) Add(template string) (int, error) {
	atomic.AddInt32(&s.addCalledCounter, 1)
	sum := 0
	delimiters, numbers := s.parseTemplate(template)
	negatives := []string{}
	for _, number := range strings.FieldsFunc(numbers, func(r rune) bool { return strings.ContainsRune(delimiters, r) }) {
		num, _ := strconv.Atoi(number)
		if num < 0 {
			negatives = append(negatives, number)
			continue
		}
		if num > 1000 {
			num = 0
		}
		sum += num
	}
	if len(negatives) != 0 {
		return 0, errors.New(fmt.Sprintf("negatives not allowed - %s", strings.Join(negatives, ",")))
	}
	return sum, nil
}

func (s *StringCalculator) parseTemplate(template string) (delimiters string, numbers string) {
	delimiters, numbers = ",\\n", template
	reg, err := regexp.Compile(`^//[\D]+\\n`)
	if err != nil {
		return
	}
	templateBytes := []byte(template)
	delimiterHeaderIndexes := reg.FindIndex(templateBytes)
	if len(delimiterHeaderIndexes) == 0 {
		return
	}
	return string(templateBytes[delimiterHeaderIndexes[0]:delimiterHeaderIndexes[1]]), string(templateBytes[delimiterHeaderIndexes[1]:])
}

func (s *StringCalculator) AddCalledCount() int {
	return int(atomic.LoadInt32(&s.addCalledCounter))
}
