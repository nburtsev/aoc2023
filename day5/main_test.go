package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var want = 35
	got := solution("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}

func TestSolution2(t *testing.T) {
	var want = 46
	got := solution2("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}

func TestSolution3(t *testing.T) {
	var want = 46
	got := solution3("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}
