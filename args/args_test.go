package args_test

import (
	"reflect"
	"testing"

	"github.com/longyue0521/TDD-In-Go/args"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	testcases := map[string]struct {
		flags    []string
		expected interface{}
	}{
		"no flags": {
			flags:    []string{},
			expected: &Option{},
		},
		"-l only": {
			flags:    []string{"-l"},
			expected: &Option{true, 0, ""},
		},
		"-p only": {
			flags:    []string{"-p", "8080"},
			expected: &Option{false, 8080, ""},
		},
		"-d only": {
			flags:    []string{"-d", "/usr/logs"},
			expected: &Option{false, 0, "/usr/logs"},
		},
		"multiple flags '-l -p 9090 -d /usr/vars'": {
			flags:    []string{"-l", "-p", "9090", "-d", "/usr/vars"},
			expected: &Option{true, 9090, "/usr/vars"},
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			// 注意:并发问题
			tt := tt
			// 利用多核,并行运行
			t.Parallel()

			actual := NewActual(tt.expected)
			args.Parse(actual, tt.flags...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func NewActual(t interface{}) interface{} {
	return reflect.New(reflect.TypeOf(t).Elem()).Interface()
}

type Option struct {
	Logging   bool   `args:"l"`
	Port      int    `args:"p"`
	Directory string `args:"d"`
}

func TestParse_should_return_error_if_tag_not_present(t *testing.T) {
	type NoTagOption struct {
		Logging   bool `args:"l"`
		Port      int  `args:"p"`
		Directory string
	}
	obj, options := &NoTagOption{}, []string{"-l", "-p", "1234", "-d", "/root/bin"}
	err := args.Parse(obj, options...)
	assert.ErrorIs(t, err, args.ErrMissingTag)
}
