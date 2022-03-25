package args

import (
	"reflect"
	"strconv"
)

func Parse(v interface{}, flags ...string) {
	obj := reflect.ValueOf(v).Elem()
	if !obj.CanSet() {
		return
	}

	var val interface{}
	if obj.Type().Field(0).IsExported() {
		val = parseBoolOption(obj, flags)
		if val != nil {
			obj.Field(0).Set(reflect.ValueOf(val))
		}
	}

	if obj.Type().Field(1).IsExported() {
		val := parseIntOption(obj, flags)
		if val != nil {
			obj.Field(1).Set(reflect.ValueOf(val))
		}
	}
	val = (interface{})(nil)
	if obj.Type().Field(2).IsExported() {
		for i, flag := range flags {
			if flag == "-d" && obj.Field(2).Type().String() == "string" {
				val = flags[i+1]
			}
		}
		if val != nil {
			obj.Field(2).Set(reflect.ValueOf(val))
		}
	}
}

func parseBoolOption(obj reflect.Value, options []string) interface{} {
	return parseOption(obj.Type().Field(0), options)
}

func parseOption(field reflect.StructField, options []string) interface{} {
	i := indexOf(options, "-"+field.Tag.Get("args"))
	var val interface{}
	if field.Type.String() == "bool" {
		if i < 0 {
			val = false
		} else {
			val = true
		}
	}
	if field.Type.String() == "int" {
		if i < 0 {
			val = 0
		} else {
			val, _ = strconv.Atoi(options[i+1])
		}
	}
	return val
}

func parseIntOption(obj reflect.Value, options []string) interface{} {
	return parseOption(obj.Type().Field(1), options)
}

func indexOf(options []string, option string) int {
	for i, opt := range options {
		if opt == option {
			return i
		}
	}
	return -1
}
