package gocap

const (
	// PERMUTATION literal
	PERMUTATION string = "permutation"

	// COMBINATION literal
	COMBINATION string = "combination"
)

// Operation model
type Operation struct {
	Type       string
	K          uint
	N          uint
	Repetition bool
}

// NewCombination constructor
func NewCombination(k uint, n uint, withRepetition bool) *Operation {
	return &Operation{
		Type:       COMBINATION,
		K:          k,
		N:          n,
		Repetition: withRepetition,
	}
}

// NewPermutation constructor
func NewPermutation(k uint, n uint, withRepetition bool) *Operation {
	return &Operation{
		Type:       PERMUTATION,
		K:          k,
		N:          n,
		Repetition: withRepetition,
	}
}
