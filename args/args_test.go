package args_test

import (
	"testing"

	"github.com/longyue0521/TDD-In-Go/args"
	"github.com/stretchr/testify/assert"
)

func TestParseOption_NoFlagsPassed_GetsDefaultValues(t *testing.T) {
	var option args.Option
	args.Parse(&option)
	assert.Equal(t, false, option.Logging())
	assert.Equal(t, 0, option.Port())
	assert.Equal(t, "", option.Directory())
}

func TestParseOption_OnlyLoggingFlagPassed_GetsTrueForLoggingAndDefaultValuesForOthers(t *testing.T) {
	var option args.Option
	args.Parse(&option, "-l")
	assert.Equal(t, true, option.Logging())
	assert.Equal(t, 0, option.Port())
	assert.Equal(t, "", option.Directory())
}
