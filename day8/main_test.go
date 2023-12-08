package main

import (
	"testing"
)


func TestSolution(t *testing.T) {

	result := solution("input_test.txt")

	if result != 6 {
		t.Errorf("Expected %d to be 6", result)
	}


}

func TestSolution2(t *testing.T) {

	result := solution2("input_test2.txt")

	if result != 6 {
		t.Errorf("Expected %d to be 6", result)
	}
}

func TestSolution3(t *testing.T) {
	result := solution3("input_test2.txt")

	if result != 6 {
		t.Errorf("Expected %d to be 6", result)
	}
}
