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
	if obj.Type().Field(0).IsExported() {
		for _, flag := range flags {
			if flag == "-l" {
				obj.Field(0).SetBool(true)
			}
		}
	}
	if obj.Type().Field(1).IsExported() {
		for i, flag := range flags {
			if flag == "-p" {
				p, _ := strconv.Atoi(flags[i+1])
				obj.Field(1).SetInt(int64(p))
			}
		}
	}
	if obj.Type().Field(2).IsExported() {
		for i, flag := range flags {
			if flag == "-d" {
				obj.Field(2).SetString(flags[i+1])
			}
		}
	}
}
