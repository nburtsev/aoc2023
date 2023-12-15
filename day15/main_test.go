package main

import (
	"testing"
)

func TestSolution1(t *testing.T) {
	want := 1320
	result := solution1("input_test.txt")
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}
}

func TestSolution2(t *testing.T) {
	want := 145
	result := solution2("input_test.txt")
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}
}

func TestHash(t *testing.T) {
	want := 52
	result := hash("HASH")
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}

	want = 30
	result = hash("rn=1")
	if result != want {
		t.Errorf("Expected %d, got %d", want, result)
	}

}
