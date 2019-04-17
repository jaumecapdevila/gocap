package gocap

// CombinationCalculator implementation
type CombinationCalculator struct{}

// Calculate the combination for the given inputs
func (c *CombinationCalculator) Calculate(operation *Operation) int {
	if operation.N == 0 && operation.R == 0 {
		return 0
	}

	if operation.Repetition {
		return c.withRepetition(operation.N, operation.R)
	}

	return c.withoutRepetition(operation.N, operation.R)
}

func (c *CombinationCalculator) withoutRepetition(n int, r int) int {
	return fact(n) / (fact(r) * fact(n-r))
}

func (c *CombinationCalculator) withRepetition(n int, r int) int {
	return fact(r+n-1) / (fact(r) * fact(n-1))
}

// NewCombinationCalculator constructor
func NewCombinationCalculator() *CombinationCalculator {
	return &CombinationCalculator{}
}
