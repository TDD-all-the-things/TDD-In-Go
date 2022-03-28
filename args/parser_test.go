package args_test

import (
	"testing"

	"github.com/longyue0521/TDD-In-Go/args"
	"github.com/stretchr/testify/assert"
)

func TestBoolOptionParser_WithExtraArgument_ReturnsError(t *testing.T) {
	options, option := []string{"-l", "t"}, "l"
	value, err := args.BoolOptionParser().Parse(options, option)
	assert.Nil(t, value)
	assert.ErrorIs(t, err, args.ErrTooManyArguments)
}

func TestBoolOptionParser_WithMoreExtraArguments_ReturnsError(t *testing.T) {
	options, option := []string{"-l", "t", "f"}, "l"
	value, err := args.BoolOptionParser().Parse(options, option)
	assert.Nil(t, value)
	assert.ErrorIs(t, err, args.ErrTooManyArguments)
}

func TestBoolOptionParser_NoFlag_ReturnsDefaultValue(t *testing.T) {
	options, option := []string{"t", "f"}, "l"
	value, err := args.BoolOptionParser().Parse(options, option)
	assert.Equal(t, false, value)
	assert.NoError(t, err)
}
