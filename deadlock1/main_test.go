package main

import "testing"

func TestProcessConcurrently(t *testing.T) {
	// arrange
	in := []int{1,2,4,3}
	exp := []bool{true, false}

	// act
	out := processConcurrently(in)

	// assert
	if len(exp) != len(out) {
		t.Error("slice length mismatch")
	}

	for i, v := range out {
		if v != exp[i] {
			t.Error("result mismatch")
		}
	}
}
