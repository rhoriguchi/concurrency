package main

import (
	"reflect"
	"testing"
)

func TestRemoveIndex_unmodified(t *testing.T) {
	// arrange
	s := []int{7, 1, 12, 3, 89, 5, 6, 13, 8, 9}
	origLen := len(s)

	// act
	s2 := removeIndex(s, 3)

	// assert
	if len(s2) != origLen-1 {
		t.Errorf("wrong length %v", len(s2))
	}

	if len(s) != origLen {
		t.Errorf("original slice length has changed to %v", len(s))
	}

	if !reflect.DeepEqual(s, []int{7, 1, 12, 3, 89, 5, 6, 13, 8, 9}) {
		t.Errorf("original slice was modified to %v\norig: %v", s, []int{7, 1, 12, 3, 89, 5, 6, 13, 8, 9})
	}
}
