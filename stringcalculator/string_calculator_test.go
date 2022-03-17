package stringcalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_SingleNumber_ReturnsThatNumbers(t *testing.T) {
	sc := NewStringCalculator()
	assert.NotNil(t, sc)
	assert.Equal(t, 1, sc.Add("1"))
}

func Test_StringCalculator(t *testing.T) {

	testcases := map[string]struct {
		numbers  string
		expected int
	}{
		"Add_EmptyString_ReturnsZero": {
			numbers:  "",
			expected: 0,
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			sc := NewStringCalculator()
			assert.NotNil(t, sc)
			assert.Equal(t, tt.expected, sc.Add(tt.numbers))
		})
	}

}
