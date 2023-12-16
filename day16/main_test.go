package main

import (
	"testing"
)

func TestSolution1(t *testing.T) {
	testCases := []struct {
		filename string
		want     int
	}{
		{"input_test.txt", 46},
	}

	for _, tc := range testCases {
		got := solution1(tc.filename)
		if got != tc.want {
			t.Errorf("solution1(%v) = %v, want %v", tc.filename, got, tc.want)
		}
	}
}

func TestSolution2(t *testing.T) {
	testCases := []struct {
		filename string
		want     int
	}{
		{"input_test.txt", 51},
	}

	for _, tc := range testCases {
		got := solution2(tc.filename)
		if got != tc.want {
			t.Errorf("solution1(%v) = %v, want %v", tc.filename, got, tc.want)
		}
	}
}
