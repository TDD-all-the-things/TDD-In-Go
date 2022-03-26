package args

import (
	"reflect"
	"strconv"
)

func Parse(v interface{}, options ...string) {
	obj := reflect.ValueOf(v).Elem()

	for i := 0; i < obj.NumField(); i++ {
		field := obj.Type().Field(i)
		if field.IsExported() {
			obj.Field(i).Set(reflect.ValueOf(parseOption(field, options)))
		}
	}
}

func parseOption(field reflect.StructField, options []string) interface{} {
	var val interface{}
	option := field.Tag.Get("args")
	val = PARSERS[field.Type.String()].Parse(options, option)
	return val
}

var PARSERS map[string]OptionParser = map[string]OptionParser{
	"bool":   BoolOptionParser(),
	"int":    SingleValueOptionParser(0, func(s string) (interface{}, error) { return strconv.Atoi(s) }),
	"string": SingleValueOptionParser("", func(s string) (interface{}, error) { return s, nil }),
}

func indexOf(options []string, option string) int {
	for i, opt := range options {
		if opt == option {
			return i
		}
	}
	return -1
}

type OptionParser interface {
	Parse(options []string, option string) interface{}
}

type boolOptionParser struct{}

func BoolOptionParser() OptionParser {
	return &boolOptionParser{}
}

func (p *boolOptionParser) Parse(options []string, option string) interface{} {
	if indexOf(options, "-"+option) < 0 {
		return false
	}
	return true
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

func (p *singleValueOptionParser) Parse(options []string, option string) interface{} {
	i := indexOf(options, "-"+option)
	if i < 0 {
		return p.defaultValue
	}
	val, _ := p.parseValueFunc(options[i+1])
	return val
}
