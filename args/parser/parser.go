package parser

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrTooManyArguments = errors.New("too many arguments")
	ErrMissingArgument  = errors.New("missing argument")
)

type OptionParser interface {
	Parse(options []string, option string) (interface{}, error)
}

type boolOptionParser struct{}

func BoolOptionParser() OptionParser {
	return &boolOptionParser{}
}

func (p *boolOptionParser) Parse(options []string, option string) (interface{}, error) {
	i := indexOf(options, "-"+option)
	if i < 0 {
		return false, nil
	}
	n := 0
	_, err := valuesOf(i+1, n, options)
	if err != nil {
		return nil, err
	}
	return true, nil
}

type singleValueOptionParser struct {
	defaultValue   interface{}
	parseValueFunc func(s string) (interface{}, error)
}

func SingleValueOptionParser(defaultValue interface{}, parseValueFunc func(s string) (interface{}, error)) OptionParser {
	return &singleValueOptionParser{
		defaultValue:   defaultValue,
		parseValueFunc: parseValueFunc,
	}
}

func (p *singleValueOptionParser) Parse(options []string, option string) (interface{}, error) {
	i := indexOf(options, "-"+option)
	if i < 0 {
		return p.defaultValue, nil
	}
	n := 1
	vals, err := valuesOf(i+1, n, options)
	if err != nil {
		return nil, err
	}
	val, _ := p.parseValueFunc(vals[0])
	return val, nil
}

func valuesOf(start int, expectedLen int, options []string) ([]string, error) {
	values := valuesOfOptionFrom(start, indexOfFirstOptionFrom(start, options), options)
	if len(values) < expectedLen {
		return nil, fmt.Errorf("%w", ErrMissingArgument)
	}
	if len(values) > expectedLen {
		return nil, fmt.Errorf("%w", ErrTooManyArguments)
	}
	return values, nil
}

func valuesOfOptionFrom(start int, end int, options []string) []string {
	values := []string{}
	for i := start; i < end; i++ {
		values = append(values, options[i])
	}
	return values
}

func indexOfFirstOptionFrom(start int, options []string) int {
	for i := start; i < len(options); i++ {
		if strings.HasPrefix(options[i], "-") {
			return i
		}
	}
	return len(options)
}

func indexOf(options []string, option string) int {
	for i, opt := range options {
		if opt == option {
			return i
		}
	}
	return -1
}
