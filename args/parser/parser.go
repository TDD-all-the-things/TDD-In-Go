package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrTooManyArguments = errors.New("too many arguments")
	ErrMissingArgument  = errors.New("missing argument")
)

type OptionParser interface {
	Parse(options []string, option string) (interface{}, error)
}

type OptionType interface {
	bool | int | string
}
type unaryOptionParser[T OptionType] struct {
	defaultValue         T
	numOfFollowingValues int
	parseVauleFunc       func(s ...string) (T, error)
}

func (p *unaryOptionParser[T]) Parse(options []string, option string) (interface{}, error) {
	i := indexOf(options, "-"+option)
	if i < 0 {
		return p.defaultValue, nil
	}
	n := p.numOfFollowingValues
	vals, err := valuesOf(i+1, n, options)
	if err != nil {
		return nil, err
	}
	return p.parseVauleFunc(vals...)
}

func BoolOptionParser() OptionParser {
	return &unaryOptionParser[bool]{
		defaultValue:         false,
		numOfFollowingValues: 0,
		parseVauleFunc: func(s ...string) (bool, error) {
			return true, nil
		},
	}
}

func IntOptionParser() OptionParser {
	return &unaryOptionParser[int]{
		defaultValue:         0,
		numOfFollowingValues: 1,
		parseVauleFunc: func(s ...string) (int, error) {
			return strconv.Atoi(s[0])
		},
	}
}

func StringOptionParser() OptionParser {
	return &unaryOptionParser[string]{
		defaultValue:         "",
		numOfFollowingValues: 1,
		parseVauleFunc: func(s ...string) (string, error) {
			return s[0], nil
		},
	}
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
