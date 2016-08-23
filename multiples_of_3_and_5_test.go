package euler

import "testing"

func TestSumOfMultiplesOf3And5Below(t *testing.T) {
	tests := []struct {
		target   int
		expected int
	}{{
		target:   10,
		expected: 23,
	}}
	for _, test := range tests {
		if sumOfMultiplesOf3And5Below(test.target) != test.expected {
			t.Fatalf("Expected to get %d. Got %v", test.expected, sumOfMultiplesOf3And5Below(test.target))
		}
	}
}
