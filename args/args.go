package args

import "strconv"

func Parse(option *Option, flags ...string) {
	for i, flag := range flags {
		if flag == "-l" {
			option.logging = true
		} else if flag == "-p" {
			option.port, _ = strconv.Atoi(flags[i+1])
		} else if flag == "-d" {
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
