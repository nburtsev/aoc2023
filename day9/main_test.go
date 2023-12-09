package main

import (
	"testing"
)


func TestSolution(t *testing.T) {

	result := solution("input_test.txt")

	if (result[0] != 114 || result[1] != 2) {
		t.Errorf("Expected %d %d to be 114 and 2", result[0], result[1])
	}


}
