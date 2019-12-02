package main

import "testing"

func TestCorrectness(t *testing.T) {
	inputs := [][]int{{
		1, 0, 0, 0, 99,
	}, {
		2, 3, 0, 3, 99,
	}, {
		2, 4, 4, 5, 99, 0,
	}, {
		1, 1, 1, 4, 99, 5, 6, 0, 99,
	}}

	outputs := [][]int{{
		2, 0, 0, 0, 99,
	}, {
		2, 3, 0, 6, 99,
	}, {
		2, 4, 4, 5, 99, 9801,
	}, {
		30, 1, 1, 4, 2, 5, 6, 0, 99,
	}}

	for i, in := range inputs {
		observed := executeInstructions(in)
		expected := outputs[i]
		match(t, expected, observed)
	}
}

func match(t *testing.T, expected, observed []int) {
	if e, o := len(expected), len(observed); e != o {
		t.Fatalf("Expected length %v. Found %v", e, o)
	}
	for i := 0; i < len(expected); i++ {
		if e, o := expected[i], observed[i]; e != o {
			t.Fatalf("Iteration %v: expected value %v, found %v", i+1, e, o)
		}
	}
}
