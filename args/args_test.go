package args_test

import (
	"reflect"
	"testing"

	"github.com/longyue0521/TDD-In-Go/args"
	"github.com/stretchr/testify/assert"
)

type Option struct {
	Logging   bool
	Port      int
	Directory string
}

func TestParseOption(t *testing.T) {

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
			expected:  &Option{false, 0, "/usr/logs"},
		},
		"multiple flags '-l -p 9090 -d /usr/vars'": {
			flags:    []string{"-l", "-p", "9090", "-d", "/usr/vars"},
			expected:  &Option{true, 9090, "/usr/vars"},
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

type AnotherOption struct {
	L bool
	P int
	D string
}

func TestParseAnotherOption(t *testing.T) {
	testcases := map[string]struct {
		flags    []string
		expected interface{}
	}{
		"no flags should get default value": {
			flags:    []string{},
			expected: &AnotherOption{},
		},
		"-l only should set true for bool field": {
			flags:    []string{"-l"},
			expected: &AnotherOption{true, 0, ""},
		},
		"-p only": {
			flags:    []string{"-p", "8080"},
			expected: &AnotherOption{false, 8080, ""},
		},
		"-d only": {
			flags:    []string{"-d", "/usr/logs"},
			expected: &AnotherOption{false, 0, "/usr/logs"},
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
