package args

func Parse(option *Option, flags ...string) {
	for _, flag := range flags {
		if flag == "-l" {
			option.logging = true
		}
	}
}

type Option struct {
	logging bool
}

func (o *Option) Logging() bool {
	return o.logging
}

func (o *Option) Port() int {
	return 0
}
func (o *Option) Directory() string {
	return ""
}
