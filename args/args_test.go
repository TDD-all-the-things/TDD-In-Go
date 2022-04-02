package args_test

import (
	"reflect"
	"testing"

	"github.com/longyue0521/TDD-In-Go/args"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	testcases := map[string]struct {
		flags     []string
		expected  interface{}
		assertion assert.ErrorAssertionFunc
	}{
		"no flags": {
			flags:     []string{},
			expected:  &Option{},
			assertion: assert.NoError,
		},
		"-l only": {
			flags:     []string{"-l"},
			expected:  &Option{true, 0, ""},
			assertion: assert.NoError,
		},
		"-p only": {
			flags:     []string{"-p", "8080"},
			expected:  &Option{false, 8080, ""},
			assertion: assert.NoError,
		},
		"-d only": {
			flags:     []string{"-d", "/usr/logs"},
			expected:  &Option{false, 0, "/usr/logs"},
			assertion: assert.NoError,
		},
		"multiple flags '-l -p 9090 -d /usr/vars'": {
			flags:     []string{"-l", "-p", "9090", "-d", "/usr/vars"},
			expected:  &Option{true, 9090, "/usr/vars"},
			assertion: assert.NoError,
		},
		"should return error if tag not present": {
			flags: []string{"-l", "-p", "9090", "-d", "/usr/vars"},
			expected: &struct {
				Logging   bool `args:"l"`
				Port      int
				Directory string `args:"d"`
			}{true, 0, ""},
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, args.ErrMissingTag)
			},
		},
		"should return error if unknown option type present": {
			flags: []string{"-k", "true", "false", "true"},
			expected: &struct {
				List []bool `args:"k"`
			}{
				[]bool(nil),
			},
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, args.ErrUnsupportedOptionType)
			},
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			// 注意:并发问题
			tt := tt
			// 利用多核,并行运行
			t.Parallel()

			actual := NewActual(tt.expected)
			err := args.Parse(actual, tt.flags...)
			tt.assertion(t, err)
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
