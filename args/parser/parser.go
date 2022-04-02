package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrTooManyArguments   = errors.New("too many arguments")
	ErrMissingArgument    = errors.New("missing argument")
	ErrAtLeastOneArgument = errors.New("at least one argument")
	ErrIllegalValue       = errors.New("illegal value")
	ErrIllegalListValues  = errors.New("illegal list values")
)

type OptionParser interface {
	Parse(options []string, option string) (interface{}, error)
}

type valueCollector interface {
	collectValues(options []string, option string) (values []string, err error)
}

type valueParser interface {
	defaults() interface{}
	parse(values ...string) (interface{}, error)
}

type optionParser struct {
	valueParser
	valueCollector
}

func (p *optionParser) Parse(options []string, option string) (interface{}, error) {
	vals, err := p.collectValues(options, option)
	if err != nil {
		return nil, err
	}
	if vals == nil {
		return p.defaults(), nil
	}
	val, err := p.parse(vals...)
	if err != nil {
		return nil, err
	}
	return val, nil
}

type fixed interface {
	bool | int | string
}
type fixedNumberValueHelper[T fixed] struct {
	defaultValue        T
	numOfExpectedValues int
	parseValue          func(s ...string) (T, error)
}

func (p *fixedNumberValueHelper[T]) collectValues(options []string, option string) ([]string, error) {
	i := indexOf(options, "-"+option)
	if i < 0 {
		return nil, nil
	}
	values := valuesOfOptionFrom(i+1, indexOfFirstOptionFrom(i+1, options), options)
	if len(values) < p.numOfExpectedValues {
		return nil, fmt.Errorf("%w", ErrMissingArgument)
	}
	if len(values) > p.numOfExpectedValues {
		return nil, fmt.Errorf("%w", ErrTooManyArguments)
	}
	return values, nil
}

func (p *fixedNumberValueHelper[T]) defaults() interface{} {
	return p.defaultValue
}

func (p *fixedNumberValueHelper[T]) parse(vals ...string) (interface{}, error) {
	val, err := p.parseValue(vals...)
	if err != nil {
		return nil, fmt.Errorf("%w", ErrIllegalValue)
	}
	return val, nil
}

type parseValueFunc[T fixed | list] func(s ...string) (T, error)

func UnaryOptionParser[T fixed](defaults T, parseValue parseValueFunc[T]) OptionParser {
	helper := &fixedNumberValueHelper[T]{
		defaultValue:        defaults,
		numOfExpectedValues: 1,
		parseValue:          parseValue,
	}
	return &optionParser{
		valueCollector: helper,
		valueParser:    helper,
	}
}

func BoolOptionParser() OptionParser {
	helper := &fixedNumberValueHelper[bool]{
		defaultValue:        false,
		numOfExpectedValues: 0,
		parseValue: func(s ...string) (bool, error) {
			return true, nil
		},
	}
	return &optionParser{
		valueCollector: helper,
		valueParser:    helper,
	}
}

func IntOptionParser() OptionParser {
	return UnaryOptionParser(0, func(s ...string) (int, error) {
		return strconv.Atoi(s[0])
	})
}

func StringOptionParser() OptionParser {
	return UnaryOptionParser("", func(s ...string) (string, error) {
		return s[0], nil
	})
}

type list interface {
	[]int | []string
}
type listValueHelper[T list] struct {
	defaultValue T
	parseValue   func(s ...string) (T, error)
}

func (p *listValueHelper[T]) collectValues(options []string, option string) ([]string, error) {
	i := indexOf(options, "-"+option)
	if i < 0 {
		return nil, nil
	}
	vals := valuesOfOptionFrom(i+1, indexOfFirstOptionFrom(i+1, options), options)
	if len(vals) < 1 {
		return nil, fmt.Errorf("%w", ErrAtLeastOneArgument)
	}
	return vals, nil
}

func (p *listValueHelper[T]) defaults() interface{} {
	return p.defaultValue
}

func (p *listValueHelper[T]) parse(vals ...string) (interface{}, error) {
	val, err := p.parseValue(vals...)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrIllegalListValues, strings.Join(vals, ","))
	}
	return val, nil
}

func ListOptionParser[T list](defaults T, parseValue parseValueFunc[T]) OptionParser {
	helper := &listValueHelper[T]{
		defaultValue: defaults,
		parseValue:   parseValue,
	}
	return &optionParser{
		valueCollector: helper,
		valueParser:    helper,
	}
}

func StringListOptionParser() OptionParser {
	return ListOptionParser([]string{}, func(s ...string) ([]string, error) {
		return s, nil
	})
}

func IntListOptionParser() OptionParser {
	return ListOptionParser([]int{}, func(s ...string) ([]int, error) {
		ints := []int{}
		for _, v := range s {
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			ints = append(ints, i)
		}
		return ints, nil
	})
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
