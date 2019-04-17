package test

import (
	"testing"

	"github.com/jaumecapdevila/gocap"
)

func TestCombination(t *testing.T) {
	operation := gocap.NewCombination(16, 3, false)

	calculator := gocap.NewCombinationCalculator()

	result, err := calculator.Calculate(operation)

	if err != nil {
		t.Error("Unable to calculate the operation...")
	}

	if result != 560 {
		t.Errorf("Combination (16, 3) should be equal to 560, %d found instead;", result)
	}
}
