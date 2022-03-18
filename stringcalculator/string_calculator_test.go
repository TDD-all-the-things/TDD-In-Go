package stringcalculator

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringCalculator_Add(t *testing.T) {

	testcases := map[string]struct {
		numbers  string
		expected int
		err      error
	}{
		"EmptyString_ReturnsZero": {
			numbers:  "",
			expected: 0,
		},
		"Single Number": {
			numbers:  "1",
			expected: 1,
		},
		"Another Single Number": {
			numbers:  "2",
			expected: 2,
		},
		"Two Numbers": {
			numbers:  "1,2",
			expected: 3,
		},
		"Another Two Numbers": {
			numbers:  "3,4",
			expected: 7,
		},
		"Unknown Amount Of Numbers": {
			numbers:  "1,3,5,7,9",
			expected: 25,
		},
		"Handle NewLine Delimiter": {
			numbers:  `1\n2,3`,
			expected: 6,
		},
		"Customize Delimiter [;]": {
			numbers:  `//;\n4;5`,
			expected: 9,
		},
		"Another Customize Delimiter [.]": {
			numbers:  `//.\n4.5.1`,
			expected: 10,
		},
		"Single Negative Number": {
			numbers: "-1",
			err:     errors.New("negatives not allowed - -1"),
		},
		"Another Single Negative Number": {
			numbers: "-2",
			err:     errors.New("negatives not allowed - -2"),
		},
		"Multiple Negative Numbers": {
			numbers: "7,-2,6,-5,",
			err:     errors.New("negatives not allowed - -2,-5"),
		},
		"Numbers bigger than 1000 should be ignored": {
			numbers:  "1000,1001,1002",
			expected: 1000,
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			sc := NewStringCalculator()
			assert.NotNil(t, sc)
			actual, err := sc.Add(tt.numbers)
			if tt.err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			}
			assert.Equal(t, tt.expected, actual)
		})
	}

}

func Test_AddCalledCount(t *testing.T) {
	t.Parallel()
	sc := NewStringCalculator()
	assert.Equal(t, 0, sc.AddCalledCount())

	t.Run("Concurrent Problem", func(t *testing.T) {
		t.Parallel()
		n := 1000
		var wg sync.WaitGroup
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				sc.Add("")
				wg.Done()
			}()
		}
		wg.Wait()
		assert.Equal(t, n, sc.AddCalledCount())
	})

}
