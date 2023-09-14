package main

import (
	"fmt"
	"testing"
)

func TestFindInXMLString_TD(t *testing.T) {
	// arrange
	type test struct {
		xml    string
		needle string
		want   bool
	}

	xmldoc := `<start serial="0xa5540012">
serial ...0012 has been issued
</start>`

	tests := []test{
		{xmldoc, "0012", true},
		{xmldoc, "serial", true},
		{xmldoc, "start", false},
		{xmldoc, "\"", false},
		{xmldoc, "/", false},
		{xmldoc, "<", false},
		{xmldoc, ">", false},
		{xmldoc, "d", true},
	}

	for _, test := range tests {
		t.Run(test.needle, func(t *testing.T) {
			// act
			res := FindInXMLString(test.xml, test.needle) > -1
			fmt.Println(res)

			// assert
			if res != test.want {
				t.Errorf("expected %t, got %t", test.want, res)
			}
		})
	}
}
