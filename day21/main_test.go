package main

import "testing"

func TestSolution1(t *testing.T) {
	want := 16
	result := solution1("input_test.txt", 64, false)
	if result != want {
		t.Errorf("Expected %d to be %v", result, want)
	}
}

func TestSolution2(t *testing.T) {
	want := 6536
	result := solution2("input_test.txt", 100)
	if result != want {
		t.Errorf("Expected %d to be %v", result, want)
	}
}
