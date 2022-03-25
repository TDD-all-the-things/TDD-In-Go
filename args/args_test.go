package args_test

import (
	"testing"

	"github.com/longyue0521/TDD-In-Go/args"
	"github.com/stretchr/testify/assert"
)

func TestParseOption(t *testing.T) {

	testcases := map[string]struct {
		flags    []string
		expected args.Option
	}{
		"no flags": {
			flags:    []string{},
			expected: args.Option{},
		},
		"-l only": {
			flags:    []string{"-l"},
			expected: args.NewOption(true, 0, ""),
		},
		"-p only": {
			flags:    []string{"-p", "8080"},
			expected: args.NewOption(false, 8080, ""),
		},
		"-d only": {
			flags:    []string{"-d", "/usr/logs"},
			expected: args.NewOption(false, 0, "/usr/logs"),
		},
		"multiple flags '-l -p 9090 -d /usr/vars'": {
			flags:    []string{"-l", "-p", "9090", "-d", "/usr/vars"},
			expected: args.NewOption(true, 9090, "/usr/vars"),
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			// 注意:并发问题
			tt := tt
			// 利用多核,并行运行
			t.Parallel()

			var actual args.Option
			args.Parse(&actual, tt.flags...)
			assert.Equal(t, tt.expected.Logging(), actual.Logging())
			assert.Equal(t, tt.expected.Port(), actual.Port())
			assert.Equal(t, tt.expected.Directory(), actual.Directory())
		})
	}
}

type AnotherOption struct {
	L bool
	P int
	D string
}

func TestParseAnotherOption_NoOption_ReturnsDefaultValue(t *testing.T) {
	var actual AnotherOption
	args.Parse(&actual)

	assert.Equal(t, false, actual.L)
	assert.Equal(t, 0, actual.P)
	assert.Equal(t, "", actual.D)
}

func TestParseAnotherOption_LOptionOnly_ReturnsTrue(t *testing.T) {
	var actual AnotherOption
	args.Parse(&actual, "-l")

	assert.Equal(t, true, actual.L)
	assert.Equal(t, 0, actual.P)
	assert.Equal(t, "", actual.D)
}
