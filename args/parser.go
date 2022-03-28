package args

import "errors"

var (
	ErrTooManyArguments = errors.New("too many arguments")
)

type OptionParser interface {
	Parse(options []string, option string) (interface{}, error)
}

type boolOptionParser struct{}

func BoolOptionParser() OptionParser {
	return &boolOptionParser{}
}

func (p *boolOptionParser) Parse(options []string, option string) (interface{}, error) {
	if indexOf(options, "-"+option) < 0 {
		return false, nil
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
	val, _ := p.parseValueFunc(options[i+1])
	return val, nil
}

func indexOf(options []string, option string) int {
	for i, opt := range options {
		if opt == option {
			return i
		}
	}
	return -1
}
