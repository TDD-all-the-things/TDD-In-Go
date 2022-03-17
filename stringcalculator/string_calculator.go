package stringcalculator

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (s *StringCalculator) Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	}
	return 1
}
