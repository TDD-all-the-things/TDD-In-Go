package args

func Parse(option *Option, flags ...string) {
}

type Option struct {
}

func (o *Option) Logging() bool {
	return false
}

func (o *Option) Port() int {
	return 0
}

func (o *Option) Directory() string {
	return ""
}
