package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	want := 374
	result := solution("input_test.txt", 2)

	t.Log("Result:", result)
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}

	want = 1030
	result = solution("input_test.txt", 10)
	t.Log("Result:", result)
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}

	want = 8410
	result = solution("input_test.txt", 100)
	t.Log("Result:", result)
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}
}

func TestDistanceBetweenPoints(t *testing.T) {
	want := 9
	result := DistanceBetweenPoints(6, 1, 11, 5)
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}

	want = 8
	result = DistanceBetweenPoints(6, 1, 5, 8)
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}

	want = 15
	result = DistanceBetweenPoints(0, 4, 10, 9)
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}
}
