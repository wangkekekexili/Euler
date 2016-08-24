package euler

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	tests := []struct {
		target   int64
		expected int64
	}{{
		target:   13195,
		expected: 29,
	}, {
		target:   600851475143,
		expected: 6857,
	}}
	for _, test := range tests {
		myAnswer := largestPrimeFactor(test.target)
		if myAnswer != test.expected {
			t.Fatalf("Expected to get %d. Got %v", test.expected, myAnswer)
		}
	}
}
