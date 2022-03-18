package args

func Parse(option *Option, flags ...string) {

}

type Option struct {
}

func (o *Option) Logging() bool {
	return true
}

func (o *Option) Port() int {
	return -1
}

func (o *Option) Directory() string {
	return "1"
}
