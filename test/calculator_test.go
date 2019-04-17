package test

import (
	"testing"

	"github.com/jaumecapdevila/gocap"
)

const miscalculation = "The %s of (%d, %d) should be %d, %d instead"

var withoutR = []struct {
	n      int
	r      int
	result int
}{
	{0, 0, 0},
	{1, 1, 1},
	{16, 3, 560},
}

var withR = []struct {
	n      int
	r      int
	result int
}{
	{0, 0, 0},
	{1, 1, 1},
	{5, 3, 35},
}

func TestCombinationWithoutRepetition(t *testing.T) {

	calculator := gocap.NewCombinationCalculator()

	for _, item := range withoutR {
		operation := gocap.NewCombination(item.n, item.r, false)
		result, _ := calculator.Calculate(operation)
		if result != item.result {
			t.Errorf(miscalculation, "Combitaion", item.n, item.r, item.result, result)
		}
	}
}

func TestCombinationWithRepetition(t *testing.T) {

	calculator := gocap.NewCombinationCalculator()

	for _, item := range withR {
		operation := gocap.NewCombination(item.n, item.r, true)
		result, _ := calculator.Calculate(operation)
		if result != item.result {
			t.Errorf(miscalculation, "Combitaion", item.n, item.r, item.result, result)
		}
	}
}
