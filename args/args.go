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
	return PARSERS[field.Type.String()].Parse(options, field.Tag.Get("args"))
}

var PARSERS map[string]OptionParser = map[string]OptionParser{
	"bool":   BoolOptionParser(),
	"int":    SingleValueOptionParser(0, func(s string) (interface{}, error) { return strconv.Atoi(s) }),
	"string": SingleValueOptionParser("", func(s string) (interface{}, error) { return s, nil }),
}
