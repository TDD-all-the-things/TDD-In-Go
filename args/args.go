package args

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/longyue0521/TDD-In-Go/args/parser"
)

var (
	ErrMissingTag = errors.New("missing tag")
)

func Parse(v interface{}, options ...string) error {
	obj := reflect.ValueOf(v).Elem()

	for i := 0; i < obj.NumField(); i++ {
		field := obj.Type().Field(i)
		if field.IsExported() {
			val, err := parseOption(field, options)
			if err != nil {
				return err
			}
			obj.Field(i).Set(reflect.ValueOf(val))
		}
	}
	return nil
}

func parseOption(field reflect.StructField, options []string) (interface{}, error) {
	p := PARSERS[field.Type.String()]
	value, err := p.Parse(options, field.Tag.Get("args"))
	return value, err
}

var PARSERS map[string]parser.OptionParser = map[string]parser.OptionParser{
	"bool":   parser.BoolOptionParser(),
	"int":    parser.SingleValueOptionParser(0, func(s string) (interface{}, error) { return strconv.Atoi(s) }),
	"string": parser.SingleValueOptionParser("", func(s string) (interface{}, error) { return s, nil }),
}
