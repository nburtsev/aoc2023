package main

import "testing"

func TestSolution1(t *testing.T) {
	want := 5
	result := solution1("input_test.txt")
	if result != want {
		t.Errorf("Expected %d to be %v", result, want)
	}
}

func TestSolution2(t *testing.T) {
	want := 7
	result := solution2("input_test.txt")
	if result != want {
		t.Errorf("Expected %d to be %v", result, want)
	}
}

func TestIsSupporting(t *testing.T) {

	a := &Brick{start: Point{0, 1, 2}, end: Point{2, 1, 2}} // block we are checking

	bricks := []*Brick{
		{start: Point{0, 1, 4}, end: Point{2, 1, 4}},
		{start: Point{0, 1, 1}, end: Point{2, 1, 1}},
		{start: Point{1, 1, 3}, end: Point{1, 2, 3}},
		{start: Point{1, 0, 3}, end: Point{1, 1, 3}},
		{start: Point{1, 0, 3}, end: Point{1, 2, 3}},
	}

	tests := []struct {
		a    *Brick
		b    *Brick
		want bool
	}{
		{a: a, b: bricks[0], want: false},
		{a: a, b: bricks[1], want: false},
		{a: a, b: bricks[2], want: true},
		{a: a, b: bricks[3], want: true},
		{a: a, b: bricks[4], want: true},
	}
	for _, tt := range tests {
		t.Run(tt.a.String(), func(t *testing.T) {
			if got := tt.a.IsSupporting(*tt.b); got != tt.want {
				t.Errorf("Brick.IsSupporting() = %v, want %v, %v", got, tt.want, tt.b)
			}
		})
	}

}

// 805 is too high
// 721 is too high
// 669 is too high
