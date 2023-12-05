package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var want = 13
	got := solution("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}

func TestSolution2(t *testing.T) {
	var want = 30
	got := solution2("input_test.txt")
	if got != want {
		t.Errorf("Failed solution = %v  instead of %v ", got, want)
	}
}

func TestProcessLine2(t *testing.T) {

	var tests = []struct {
		input string
		want  struct{ x, y int }
	}{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", struct{ x, y int }{1, 4}},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", struct{ x, y int }{2, 2}},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", struct{ x, y int }{3, 2}},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", struct{ x, y int }{4, 1}},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", struct{ x, y int }{5, 0}},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", struct{ x, y int }{6, 0}},
	}

	for _, test := range tests {
		var got struct{ x, y int }

		r1, r2 := processLine2(test.input)
		if r1 != test.want.x || r2 != test.want.y {
			t.Errorf("Failed nextToSymbol(%q) = %v instead of %v", test.input, got, test.want)
		}
	}

}

// func TestSolution2(t *testing.T) {
// 	var want = 467835
// 	got := solution2("input_test.txt")
// 	if got != want {
// 		t.Errorf("Failed solution = %v  instead of %v ", got, want)
// 	}
// }

// type Test struct {
// 	number []int
// 	symbol int
// }

// func TestNexToSymbol(t *testing.T) {
// 	var tests = []struct {
// 		input Test
// 		want  int
// 	}{
// 		{Test{[]int{805, 46, 49}, 47}, 805},
// 		{Test{[]int{426, 68, 71}, 72}, 0},
// 	}

// 	for _, test := range tests {
// 		if got := nextToSymbolOrZero(test.input.number, test.input.symbol); got != test.want {
// 			t.Errorf("Failed nextToSymbol(%q) = %v instead of %v", test.input, got, test.want)
// 		}
// 	}

// }
