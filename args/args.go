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
		val = false
		for _, flag := range flags {
			if flag == "-l" && obj.Field(0).Type().String() == "bool" {
				val = true
			}
		}
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
