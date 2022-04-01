package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrTooManyArguments   = errors.New("too many arguments")
	ErrMissingArgument    = errors.New("missing argument")
	ErrAtLeastOneArgument = errors.New("at least one argument")
	ErrIllegalValue       = errors.New("illegal value")
)

type OptionParser interface {
	Parse(options []string, option string) (interface{}, error)
}

type unary interface {
	bool | int | string
}
type unaryOptionParser[T unary] struct {
	defaultValue        T
	numOfExpectedValues int
	parseValue          func(s ...string) (T, error)
}

func (p *unaryOptionParser[T]) Parse(options []string, option string) (interface{}, error) {
	i := indexOf(options, "-"+option)
	if i < 0 {
		return p.defaultValue, nil
	}
	vals, err := valuesOf(i+1, p.numOfExpectedValues, options)
	if err != nil {
		return nil, err
	}
	val, err := p.parseValue(vals...)
	if err != nil {
		return nil, fmt.Errorf("%w", ErrIllegalValue)
	}
	return val, nil
}

func BoolOptionParser() OptionParser {
	return &unaryOptionParser[bool]{
		defaultValue:        false,
		numOfExpectedValues: 0,
		parseValue: func(s ...string) (bool, error) {
			return true, nil
		},
	}
}

func IntOptionParser() OptionParser {
	return &unaryOptionParser[int]{
		defaultValue:        0,
		numOfExpectedValues: 1,
		parseValue: func(s ...string) (int, error) {
			return strconv.Atoi(s[0])
		},
	}
}

func StringOptionParser() OptionParser {
	return &unaryOptionParser[string]{
		defaultValue:        "",
		numOfExpectedValues: 1,
		parseValue: func(s ...string) (string, error) {
			return s[0], nil
		},
	}
}

type list interface {
	[]string
}
type listOptionParser[T list] struct {
	defaultValues T
	parseValues   func(s ...string) (T, error)
}

func (p *listOptionParser[T]) Parse(options []string, option string) (interface{}, error) {
	i := indexOf(options, "-"+option)
	if i < 0 {
		return p.defaultValues, nil
	}
	start := i + 1
	vals := valuesOfOptionFrom(start, indexOfFirstOptionFrom(start, options), options)
	if len(vals) < 1 {
		return nil, fmt.Errorf("%w", ErrAtLeastOneArgument)
	}
	val, _ := p.parseValues(vals...)
	return val, nil
}

func StringListParser(parseValues ...ParseValueFunc[[]string]) OptionParser {
	parser := &listOptionParser[[]string]{
		defaultValues: []string{},
		parseValues: func(s ...string) ([]string, error) {
			return s, nil
		},
	}
	if len(parseValues) != 0 {
		parser.parseValues = parseValues[0]
	}
	return parser
}

type ParseValueFunc[T any] func(s ...string) (T, error)

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
		if isOption(options[i]) {
			return i
		}
	}
	return len(options)
}

func isOption(option string) bool {
	ok, _ := regexp.MatchString(`^-[a-zA-Z-]+$`, option)
	return ok
}

func indexOf(options []string, option string) int {
	for i, opt := range options {
		if opt == option {
			return i
		}
	}
	return -1
}
