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
