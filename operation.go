package gocap

// Calculator defines the contract to calculate an operation
type Calculator interface {
	Calculate(*Operation) int
}

// Operation model
type Operation struct {
	N          int
	R          int
	Repetition bool
}

// NewOperation constructor
func NewOperation(n int, r int, withRepetition bool) *Operation {
	return &Operation{
		N:          n,
		R:          r,
		Repetition: withRepetition,
	}
}
