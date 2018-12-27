package aoc

import "testing"

func TestMin(t *testing.T) {
	tt := []struct {
		name string
		nums []int
		min  int
	}{
		{"just one", []int{1}, 1},
		{"one and two", []int{2, 1}, 1},
		{"one and one", []int{1, 1}, 1},
		{"empty", nil, maxInt},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if min := Min(tc.nums...); min != tc.min {
				t.Fatalf("expected min of %v to be %v; got %v", tc.nums, tc.min, min)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tt := []struct {
		name string
		nums []int
		max  int
	}{
		{"just one", []int{1}, 1},
		{"one and two", []int{2, 1}, 2},
		{"one and one", []int{1, 1}, 1},
		{"empty", nil, minInt},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if max := Max(tc.nums...); max != tc.max {
				t.Fatalf("expected max of %v to be %v; got %v", tc.nums, tc.max, max)
			}
		})
	}
}

func TestMinArg(t *testing.T) {
	tt := []struct {
		name string
		nums []int
		min  int
		arg  int
	}{
		{"just one", []int{1}, 1, 0},
		{"one and two", []int{2, 1}, 1, 1},
		{"one and one", []int{1, 1}, 1, 0},
		{"empty", nil, maxInt, -1},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if min := Min(tc.nums...); min != tc.min {
				t.Fatalf("expected min of %v to be %v; got %v", tc.nums, tc.min, min)
			}
		})
	}
}

func TestMaxArg(t *testing.T) {
	tt := []struct {
		name string
		nums []int
		max  int
		arg  int
	}{
		{"just one", []int{1}, 1, 0},
		{"one and two", []int{2, 1}, 2, 0},
		{"one and one", []int{1, 1}, 1, 0},
		{"empty", nil, minInt, -1},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if max := Max(tc.nums...); max != tc.max {
				t.Fatalf("expected max of %v to be %v; got %v", tc.nums, tc.max, max)
			}
		})
	}
}
