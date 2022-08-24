package main

import "testing"

func TestSum(t *testing.T) {
	e1 := elements{1, 3, 5, 7}
	expectedSum := 16
	sum1 := sumElements(e1)

	if sum1 != expectedSum {
		t.Errorf("Expected sum %d, but got %d", expectedSum, sum1)

	}

}

func TestMin(t *testing.T) {
	e1 := elements{34, 12, 3, 1, 5, 7}
	expectedSum := 1
	sum1 := minElement(e1)

	if sum1 != expectedSum {
		t.Errorf("Expected sum %d, but got %d", expectedSum, sum1)

	}

}
