package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	want := 21
	result := solution("input_test.txt")
	if result != want {
		t.Errorf("Expected result to be %d, got %d", want, result)
	}

}

func TestSolution2(t *testing.T) {
	want := 525152
	result := solution2("input_test.txt")
	if result != want {
		t.Errorf("Expected result to be %d, got %d", want, result)
	}

}
