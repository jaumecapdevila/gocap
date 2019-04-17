package gocap

import "fmt"

// CombinationCalculator implementation
type CombinationCalculator struct{}

// Calculate the combination for the given inputs
func (c *CombinationCalculator) Calculate(operation *Operation) (int, error) {
	if operation.Type != COMBINATION {
		return 0, fmt.Errorf("Invalid operation")
	}

	if operation.Repetition {
		return withRepetition(operation.N, operation.R), nil
	}

	return withoutRepetition(operation.N, operation.R), nil
}

func withoutRepetition(n int, r int) int {
	return fact(n) / (fact(r) * fact(n-r))
}

func withRepetition(n int, r int) int {
	return fact(r+n-1) / (fact(r) * fact(n-1))
}

// NewCombinationCalculator constructor
func NewCombinationCalculator() *CombinationCalculator {
	return &CombinationCalculator{}
}
