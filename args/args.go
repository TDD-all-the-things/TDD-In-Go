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

	val = (interface{})(nil)
	if obj.Type().Field(1).IsExported() {
		for i, flag := range flags {
			if flag == "-p" && obj.Field(1).Type().String() == "int" {
				val, _ = strconv.Atoi(flags[i+1])
			}
		}
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
	i := indexOf(options, "-l")
	var val interface{}
	if obj.Field(0).Type().String() == "bool" {
		if i < 0 {
			val = false
		} else {
			val = true
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
