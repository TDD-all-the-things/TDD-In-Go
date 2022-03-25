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
	for i := 0; i < 2; i++ {
		if obj.Type().Field(i).IsExported() {
			obj.Field(i).Set(reflect.ValueOf(parseOption(obj.Type().Field(i), flags)))
		}
	}

	if obj.Type().Field(2).IsExported() {
		obj.Field(2).Set(reflect.ValueOf(parseOption(obj.Type().Field(2), flags)))
	}
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
	if field.Type.String() == "string" {
		if i < 0 {
			val = ""
		} else {
			val = options[i+1]
		}
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
