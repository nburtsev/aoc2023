package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var want = 288
	got := solution("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}

func TestSolution2(t *testing.T) {
	var want = 71503
	got := solution2("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}
