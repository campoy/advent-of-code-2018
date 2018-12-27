package main

import (
	"fmt"
	"testing"
)

func TestScore(t *testing.T) {
	tt := []struct {
		players int
		last    int
		score   int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d players %d last", tc.players, tc.last), func(t *testing.T) {
			if s := score(tc.players, tc.last, nil); s != tc.score {
				t.Fatalf("expected score to be %d; got %d", tc.score, s)
			}
		})
	}
}
