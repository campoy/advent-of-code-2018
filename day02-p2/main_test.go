package main

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	tt := []struct {
		a, b   string
		common string
		ok     bool
	}{
		{"abcde", "fghij", "", false},
		{"abcde", "abcue", "abce", true},
		{"abcde", "abcdu", "abcd", true},
		{"abcde", "ubcde", "bcde", true},
		{"abcde", "uddbcde", "", false},
		{"abcde", "", "", false},
		{"abcde", "abcde", "", false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s vs %s", tc.a, tc.b), func(t *testing.T) {
			c, ok := compare(tc.a, tc.b)
			if tc.ok != ok {
				t.Fatalf("expected ok to be %v; got %v", tc.ok, ok)
			}
			if tc.common != c {
				t.Fatalf("expected common to be %s; got %s", tc.common, c)
			}
		})
	}
}
