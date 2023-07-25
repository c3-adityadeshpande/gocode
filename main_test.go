// main_test.go
package main

import (
	"fmt"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{5, 10, 15},
		{0, 0, 0},
		{-5, 5, 0},
		{-10, -20, -30},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("AddNumbers(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			result := AddNumbers(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("AddNumbers(%d, %d) = %d, want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
