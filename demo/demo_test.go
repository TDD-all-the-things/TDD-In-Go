package demo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Demo(t *testing.T) {
	assert.Equal(t, "Hello, World!", "Hello, Go")
}
