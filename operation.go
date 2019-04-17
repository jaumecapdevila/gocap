package gocap

const (
	// PERMUTATION literal
	PERMUTATION string = "permutation"

	// COMBINATION literal
	COMBINATION string = "combination"
)

// Calculator defines the contract to calculate an operation
type Calculator interface {
	Calculate(*Operation) (int, error)
}

// Operation model
type Operation struct {
	Type       string
	N          int
	R          int
	Repetition bool
}

// NewCombination constructor
func NewCombination(n int, r int, withRepetition bool) *Operation {
	return &Operation{
		Type:       COMBINATION,
		N:          n,
		R:          r,
		Repetition: withRepetition,
	}
}

// NewPermutation constructor
func NewPermutation(n int, r int, withRepetition bool) *Operation {
	return &Operation{
		Type:       PERMUTATION,
		N:          n,
		R:          r,
		Repetition: withRepetition,
	}
}
