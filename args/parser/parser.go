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
	values(options []string, option string) (values []string, err error)
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
	vals, err := p.values(options, option)
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

type unary interface {
	bool | int | string
}
type unaryOptionParser[T unary] struct {
	defaultValue        T
	numOfExpectedValues int
	parseValue          func(s ...string) (T, error)
}

func (p *unaryOptionParser[T]) values(options []string, option string) ([]string, error) {
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

func (p *unaryOptionParser[T]) defaults() interface{} {
	return p.defaultValue
}

func (p *unaryOptionParser[T]) parse(vals ...string) (interface{}, error) {
	val, err := p.parseValue(vals...)
	if err != nil {
		return nil, fmt.Errorf("%w", ErrIllegalValue)
	}
	return val, nil
}

type parseValueFunc[T unary | list] func(s ...string) (T, error)

func UnaryOptionParser[T unary](defaults T, numOfExpectedValues int, parseValue parseValueFunc[T]) OptionParser {
	valuer := &unaryOptionParser[T]{
		defaultValue:        defaults,
		numOfExpectedValues: numOfExpectedValues,
		parseValue:          parseValue,
	}
	return &optionParser{
		valueCollector: valuer,
		valueParser:    valuer,
	}
}

func BoolOptionParser() OptionParser {
	return UnaryOptionParser(false, 0, func(s ...string) (bool, error) {
		return true, nil
	})
}

func IntOptionParser() OptionParser {
	return UnaryOptionParser(0, 1, func(s ...string) (int, error) {
		return strconv.Atoi(s[0])
	})
}

func StringOptionParser() OptionParser {
	return UnaryOptionParser("", 1, func(s ...string) (string, error) {
		return s[0], nil
	})
}

type list interface {
	[]string
}
type listOptionParser[T list] struct {
	defaultValue T
	parseValue   func(s ...string) (T, error)
}

func (p *listOptionParser[T]) values(options []string, option string) ([]string, error) {
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

func (p *listOptionParser[T]) defaults() interface{} {
	return p.defaultValue
}

func (p *listOptionParser[T]) parse(vals ...string) (interface{}, error) {
	val, err := p.parseValue(vals...)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrIllegalListValues, strings.Join(vals, ","))
	}
	return val, nil
}

func ListOptionParser[T list](defaults T, parseValue parseValueFunc[T]) OptionParser {
	valuer := &listOptionParser[T]{
		defaultValue: defaults,
		parseValue:   parseValue,
	}
	return &optionParser{
		valueCollector: valuer,
		valueParser:    valuer,
	}
}

func StringListOptionParser(parseValue parseValueFunc[[]string]) OptionParser {
	return ListOptionParser([]string{}, parseValue)
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
