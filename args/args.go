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
	if field.Type.String() == "bool" {
		option := field.Tag.Get("args")
		val = parseBoolOption(options, option)
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
	i := indexOf(options, "-"+option)
	if i < 0 {
		val = ""
	} else {
		val = options[i+1]
	}
	return val
}

func parseIntOption(options []string, option string) interface{} {
	var val interface{}
	i := indexOf(options, "-"+option)
	if i < 0 {
		val = 0
	} else {
		val, _ = strconv.Atoi(options[i+1])
	}
	return val
}

func parseBoolOption(options []string, option string) interface{} {
	var val interface{}
	i := indexOf(options, "-"+option)
	if i < 0 {
		val = false
	} else {
		val = true
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
