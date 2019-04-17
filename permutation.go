package gocap

import (
	"fmt"
	"math"
)

// PermutationCalculator implementation
type PermutationCalculator struct{}

// Calculate the permutation for the given inputs
func (c *PermutationCalculator) Calculate(operation *Operation) (int, error) {
	if operation.Type != PERMUTATION {
		return 0, fmt.Errorf("Invalid operation")
	}

	if operation.N == 0 && operation.R == 0 {
		return 0, nil
	}

	if operation.Repetition {
		return c.withRepetition(operation.N, operation.R), nil
	}

	return c.withoutRepetition(operation.N, operation.R), nil
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
