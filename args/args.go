package args

import (
	"reflect"
	"strconv"
)

func Parse(v interface{}, flags ...string) {
	option, _ := v.(*Option)
	obj := reflect.ValueOf(v).Elem()
	for i, flag := range flags {
		if flag == "-l" {
			if obj.CanSet() && obj.Type().Field(0).IsExported() {
				obj.Field(0).SetBool(true)
				continue
			}
			option.logging = true
		} else if flag == "-p" {
			if obj.CanSet() && obj.Type().Field(1).IsExported() {
				p, _ := strconv.Atoi(flags[i+1])
				obj.Field(1).SetInt(int64(p))
				continue
			}
			option.port, _ = strconv.Atoi(flags[i+1])
		} else if flag == "-d" {
			if obj.CanSet() && obj.Type().Field(2).IsExported() {
				obj.Field(2).SetString(flags[i+1])
				continue
			}
			option.directory = flags[i+1]
		}
	}
}

type Option struct {
	logging   bool
	port      int
	directory string
}

func NewOption(logging bool, port int, directory string) Option {
	return Option{logging: logging, port: port, directory: directory}
}

func (o *Option) Logging() bool {
	return o.logging
}

func (o *Option) Port() int {
	return o.port
}

func (o *Option) Directory() string {
	return o.directory
}
