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
		"SingleNumber_ReturnsThatNumber": {
			numbers:  "1",
			expected: 1,
		},
		"AnotherSingleNumber_ReturnsThatNumber": {
			numbers:  "2",
			expected: 2,
		},
		"TwoNumbers_ReturnsSumOfBoth": {
			numbers:  "1,2",
			expected: 3,
		},
		"AnotherTwoNumbers_ReturnsSumOfBoth": {
			numbers:  "3,4",
			expected: 7,
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

func Test_Add_UnknownAmountOfNumbers_ReturnsSumOfThem(t *testing.T) {
	sc := NewStringCalculator()
	assert.Equal(t, 25, sc.Add("1,3,5,7,9"))
}
