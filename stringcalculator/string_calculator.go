package stringcalculator

import (
	"bytes"
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
	delimiter, numbers := s.parseTemplate(template)
	negatives := []string{}
	for _, number := range strings.Split(strings.ReplaceAll(numbers, `\n`, delimiter), delimiter) {
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
		return 0, errors.New(fmt.Sprintf("negatives not allowed - %s", strings.Join(negatives, delimiter)))
	}
	return sum, nil
}

func (s *StringCalculator) parseTemplate(template string) (delimiter string, numbers string) {
	delimiter, numbers = ",", template
	reg, err := regexp.Compile(`^//[\D]+\\n`)
	if err != nil {
		return
	}
	templateBytes := []byte(template)
	delimiterHeaderIndexes := reg.FindIndex(templateBytes)
	if len(delimiterHeaderIndexes) == 0 {
		return
	}
	return s.parseTemplateBytes(delimiterHeaderIndexes, templateBytes)
}

func (s *StringCalculator) parseTemplateBytes(delimiterHeaderIndexes []int, templateBytes []byte) (string, string) {
	delimiterHeaderStartIndex, delimiterHeaderEndIndex := delimiterHeaderIndexes[0], delimiterHeaderIndexes[1]
	delimiterStartIndex, delimiterEndIndex := delimiterHeaderStartIndex+len(`//`), delimiterHeaderEndIndex-len(`\n`)
	if bytes.HasPrefix(templateBytes[delimiterStartIndex:delimiterEndIndex], []byte("[")) {
		delimiterStartIndex += len("[")
		delimiterEndIndex -= len("]")
	}
	return string(templateBytes[delimiterStartIndex:delimiterEndIndex]), string(templateBytes[delimiterHeaderEndIndex:])
}

func (s *StringCalculator) AddCalledCount() int {
	return int(atomic.LoadInt32(&s.addCalledCounter))
}
