package main

import (
	"fmt"
	"testing"
)

func TestPowerAt(t *testing.T) {
	tt := []struct {
		x, y   int
		serial int
		power  int
	}{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			p := powerAt(tc.x, tc.y, tc.serial)
			if p != tc.power {
				t.Fatalf("expected power to be %d; got %d", tc.power, p)
			}
		})
	}
}
