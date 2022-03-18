package args

func Parse(option *Option, flags ...string) {
	for _, flag := range flags {
		if flag == "-l" {
			option.logging = true
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
