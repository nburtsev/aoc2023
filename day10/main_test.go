package main

import (
	"testing"
)

func TestSolution(t *testing.T) {

	want := 8
	result := solution("input_test.txt")

	if result != want {
		t.Errorf("Expected %d to be %d", result, want)
	}

}

func TestSolution2(t *testing.T) {

	want := 3
	result := solution2("input_test2.txt")

	if result != want {
		t.Errorf("Expected %d to be %d", result, want)
	}

	want = 9
	result = solution2("input_test4.txt")

	if result != want {
		t.Errorf("Expected %d to be %d", result, want)
	}

	want = 8
	result = solution2("input_test3.txt")

	if result != want {
		t.Errorf("Expected %d to be %d", result, want)
	}

	want = 10
	result = solution2("input_test4.txt")

	if result != want {
		t.Errorf("Expected %d to be %d", result, want)
	}

	want = 493
	result = solution2("input.txt")

	if result >= 493 || result <= 429 {
		t.Errorf("wrong answer, must be larger than 429 and smaller than 493")
	}

}
