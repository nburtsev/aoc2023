package main

import (
	"testing"
)

func TestProcess(t *testing.T) {
	var tests = []struct {
		input string
		want  GameResult
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", GameResult{1, true, 48}},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", GameResult{2, true, 12}},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", GameResult{3, false, 1560}},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", GameResult{4, false, 630}},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", GameResult{5, true, 36}},
	}
	for _, test := range tests {
		if got := processLine(test.input, CubeHand{12, 13, 14}); got != test.want {
			t.Errorf("Failed process(%q) = %v", test.input, got)
		}
	}
}
