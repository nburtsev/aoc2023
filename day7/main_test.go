package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var want = 253910319
	got := solution("input.txt", "23456789TJQKA", handStrength)
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}

func TestSolution2(t *testing.T) {
	var want = 254083736
	got := solution("input.txt", "J23456789TQKA", handStrength2)
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}
