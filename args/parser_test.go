package args_test

import (
	"testing"

	"github.com/longyue0521/TDD-In-Go/args"
	"github.com/stretchr/testify/assert"
)

func TestBoolOptionParser(t *testing.T) {
	testcases := map[string]struct {
		options   []string
		option    string
		expected  interface{}
		assertion assert.ErrorAssertionFunc
	}{
		"should not accept extra argument for bool option": {
			options:  []string{"-l", "t"},
			option:   "l",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, args.ErrTooManyArguments)
			},
		},
		"should not accept more extra arguments for bool option": {
			options:  []string{"-l", "t", "f"},
			option:   "l",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, args.ErrTooManyArguments)
			},
		},
		"should get default value if bool option not present": {
			options:  []string{},
			option:   "l",
			expected: (interface{})(false),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			actual, err := args.BoolOptionParser().Parse(tt.options, tt.option)
			assert.Equal(t, tt.expected, actual)
			tt.assertion(t, err)
		})
	}
}
