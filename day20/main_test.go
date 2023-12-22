package main

import "testing"

func TestSolution1(t *testing.T) {
	want := 32000000
	result := solution1("input_test.txt")
	if result != want {
		t.Errorf("Expected %d to be %v", result, want)
	}
}

func TestSolution1_2(t *testing.T) {
	want := 11687500
	result := solution1("input_test2.txt")
	if result != want {
		t.Errorf("Expected %d to be %v", result, want)
	}
}

func TestSolution2(t *testing.T) {
	want := 0
	result := solution2("input_test.txt")
	if result != want {
		t.Errorf("Expected %d to be %v", result, want)
	}
}
