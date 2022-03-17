package stringcalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringCalculator_Add(t *testing.T) {

	testcases := map[string]struct {
		numbers  string
		expected int
	}{
		"EmptyString_ReturnsZero": {
			numbers:  "",
			expected: 0,
		},
		"SingleNumber_ReturnsThatNumbers": {
			numbers:  "1",
			expected: 1,
		},
		"AnotherSingleNumber_ReturnsThatNumbers": {
			numbers:  "2",
			expected: 2,
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
