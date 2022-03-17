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
		"SingleNumber": {
			numbers:  "1",
			expected: 1,
		},
		"AnotherSingleNumbe": {
			numbers:  "2",
			expected: 2,
		},
		"TwoNumbers": {
			numbers:  "1,2",
			expected: 3,
		},
		"AnotherTwoNumbers": {
			numbers:  "3,4",
			expected: 7,
		},
		"UnknownAmountOfNumbers": {
			numbers:  "1,3,5,7,9",
			expected: 25,
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			sc := NewStringCalculator()
			assert.NotNil(t, sc)
			assert.Equal(t, tt.expected, sc.Add(tt.numbers))
		})
	}

}
