package gocap

import (
	"math"
)

// PermutationCalculator implementation
type PermutationCalculator struct{}

// Calculate the permutation for the given inputs
func (c *PermutationCalculator) Calculate(operation *Operation) int {
	if operation.N == 0 && operation.R == 0 {
		return 0
	}

	if operation.Repetition {
		return c.withRepetition(operation.N, operation.R)
	}

	return c.withoutRepetition(operation.N, operation.R)
}

func (c *PermutationCalculator) withoutRepetition(n int, r int) int {
	return fact(n) / fact(n-r)
}

func (c *PermutationCalculator) withRepetition(n int, r int) int {
	result := math.Pow(float64(n), float64(r))
	return int(result)
}

// NewPermutationCalculator constructor
func NewPermutationCalculator() *PermutationCalculator {
	return &PermutationCalculator{}
}
