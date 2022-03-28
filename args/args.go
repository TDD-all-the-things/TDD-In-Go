package args

import (
	"reflect"
	"strconv"

	"github.com/longyue0521/TDD-In-Go/args/parser"
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
	p := PARSERS[field.Type.String()]
	value, _ := p.Parse(options, field.Tag.Get("args"))
	return value
}

var PARSERS map[string]parser.OptionParser = map[string]parser.OptionParser{
	"bool":   parser.BoolOptionParser(),
	"int":    parser.SingleValueOptionParser(0, func(s string) (interface{}, error) { return strconv.Atoi(s) }),
	"string": parser.SingleValueOptionParser("", func(s string) (interface{}, error) { return s, nil }),
}
