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
