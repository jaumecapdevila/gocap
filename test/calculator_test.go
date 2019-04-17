package test

import (
	"testing"

	"github.com/jaumecapdevila/gocap"
)

const miscalculation = "The %s of (%d, %d) should be %d, %d instead"

var combinations = []struct {
	n      int
	r      int
	result int
}{
	{0, 0, 0},
	{1, 1, 1},
	{16, 3, 560},
}

var combinationsR = []struct {
	n      int
	r      int
	result int
}{
	{0, 0, 0},
	{1, 1, 1},
	{5, 3, 35},
}

var permutations = []struct {
	n      int
	r      int
	result int
}{
	{0, 0, 0},
	{1, 1, 1},
	{16, 3, 3360},
}

var permutationsR = []struct {
	n      int
	r      int
	result int
}{
	{0, 0, 0},
	{1, 1, 1},
	{10, 3, 1000},
}

func TestCombinationWithoutRepetition(t *testing.T) {

	calculator := gocap.NewCombinationCalculator()

	for _, item := range combinations {
		operation := gocap.NewOperation(item.n, item.r, false)
		result := calculator.Calculate(operation)
		if result != item.result {
			t.Errorf(miscalculation, "Combitaion", item.n, item.r, item.result, result)
		}
	}
}

func TestCombinationWithRepetition(t *testing.T) {

	calculator := gocap.NewCombinationCalculator()

	for _, item := range combinationsR {
		operation := gocap.NewOperation(item.n, item.r, true)
		result := calculator.Calculate(operation)
		if result != item.result {
			t.Errorf(miscalculation, "Combitaion", item.n, item.r, item.result, result)
		}
	}
}

func TestPermutationWithoutRepetition(t *testing.T) {

	calculator := gocap.NewPermutationCalculator()

	for _, item := range permutations {
		operation := gocap.NewOperation(item.n, item.r, false)
		result := calculator.Calculate(operation)
		if result != item.result {
			t.Errorf(miscalculation, "Permutation", item.n, item.r, item.result, result)
		}
	}
}

func TestPermutationWithRepetition(t *testing.T) {

	calculator := gocap.NewPermutationCalculator()

	for _, item := range permutationsR {
		operation := gocap.NewOperation(item.n, item.r, true)
		result := calculator.Calculate(operation)
		if result != item.result {
			t.Errorf(miscalculation, "Permutation", item.n, item.r, item.result, result)
		}
	}
}
