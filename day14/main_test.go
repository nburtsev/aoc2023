package main

import (
	"testing"
)

func TestSolution1(t *testing.T) {
	want := 136
	result := solution1("input_test.txt")
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}
}

func TestSolution2(t *testing.T) {
	want := 64
	result := solution2("input_test.txt", TIMES)
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}
}
