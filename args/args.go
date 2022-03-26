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
	var parser OptionParser

	if field.Type.String() == "bool" {
		option := field.Tag.Get("args")
		parser = BoolOptionParser()
		val = parser.Parse(options, option)
	}
	if field.Type.String() == "int" {
		option := field.Tag.Get("args")
		val = parseIntOption(options, option)
	}
	if field.Type.String() == "string" {
		option := field.Tag.Get("args")
		val = parseStringOption(options, option)
	}
	return val
}

func parseStringOption(options []string, option string) interface{} {
	var val interface{}
	defaultValue := ""
	parseValue := func (s string) (interface{}, error) {
		return s, nil
	}
	i := indexOf(options, "-"+option)
	if i < 0 {
		val = defaultValue
	} else {
		val, _ = parseValue(options[i+1])
	}
	return val
}

func parseIntOption(options []string, option string) interface{} {
	var val interface{}
	defaultValue := 0
	parseValue := func (s string) (interface{}, error) {
		return strconv.Atoi(s)
	}
	i := indexOf(options, "-"+option)
	if i < 0 {
		val = defaultValue
	} else {
		val, _ = parseValue(options[i+1])
	}
	return val
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

type boolOptionParser struct {}

func BoolOptionParser() OptionParser {
	return &boolOptionParser{}
}

func (p *boolOptionParser) Parse(options []string, option string) interface{} {
	if indexOf(options, "-"+option) < 0 {
		return false
	}
	return true
}