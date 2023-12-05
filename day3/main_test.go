package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var want = 4361
	got := solution("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}

func TestSolution2(t *testing.T) {
	var want = 467835
	got := solution2("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}

type Test struct {
	number []int
	symbol int
}

func TestNexToSymbol(t *testing.T) {
	var tests = []struct {
		input Test
		want  int
	}{
		{Test{[]int{805, 46, 49}, 47}, 805},
		{Test{[]int{426, 68, 71}, 72}, 0},
	}

	for _, test := range tests {
		if got := nextToSymbolOrZero(test.input.number, test.input.symbol); got != test.want {
			t.Errorf("Failed nextToSymbol(%q) = %v instead of %v", test.input, got, test.want)
		}
	}

}
