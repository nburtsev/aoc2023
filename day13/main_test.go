package main

import (
	"testing"
)

func TestSolution(t *testing.T) {

	want := 405
	result := solution("input_test.txt")
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}
}

func TestSolution2(t *testing.T) {

	want := 400
	result := solution2("input_test.txt")
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}
}
