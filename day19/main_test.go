package main

import "testing"

func TestSolution1(t *testing.T) {
	result := solution1("input_test.txt")
	if result != 19114 {
		t.Errorf("Expected %d to be 19114", result)
	}
}

func TestSolution2(t *testing.T) {
	result := solution2("input_test.txt")
	if result != 167409079868000 {

		t.Errorf("Expected %d to be 167409079868000", result)
	}
}

// 41213785377150
// 41256000000000
// 167409079868000
