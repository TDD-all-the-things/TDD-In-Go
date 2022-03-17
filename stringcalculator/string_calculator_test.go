package stringcalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_EmptyString_ReturnsZero(t *testing.T) {
	sc := NewStringCalculator()
	assert.NotNil(t, sc)
	assert.Equal(t, 0, sc.Add(""))
}
