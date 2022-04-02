package args

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/longyue0521/TDD-In-Go/args/parser"
)

var (
	ErrMissingTag            = errors.New("missing tag")
	ErrUnsupportedOptionType = errors.New("unsupported option type")
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
	tag := field.Tag.Get("args")
	if tag == "" {
		return nil, fmt.Errorf("%w", ErrMissingTag)
	}
	p, ok := parsers[field.Type.String()]
	if !ok {
		return nil, fmt.Errorf("%w", ErrUnsupportedOptionType)
	}
	return p.Parse(options, tag)
}

var parsers map[string]parser.OptionParser = map[string]parser.OptionParser{
	"bool":   parser.BoolOptionParser(),
	"int":    parser.IntOptionParser(),
	"string": parser.StringOptionParser(),
}
