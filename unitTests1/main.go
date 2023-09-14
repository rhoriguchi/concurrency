package main

import (
	"fmt"
	"strings"
)

// FindInXMLString searches an XML document's content for a substring.
// Element names and attribute names will be ignored.
/*
BUGS:
 - r is 0, 1, ..., len(xml) => jump directly to last "else if" => ignores above requirement
 - if above issue is fixed, still run into problem that it simply checks inclusion in 0xa5540012">
serial ...0012 has been issued
</start>
*/
func FindInXMLString(xml string, needle string) int {
	var inElementTag bool
	var inAttributeValue bool

	var i int
	for _, r := range xml {
		if r == '<' && !inElementTag && !inAttributeValue {
			inElementTag = true
		} else if r == '>' && inElementTag && !inAttributeValue {
			inElementTag = false
		} else if r == '"' && inElementTag && !inAttributeValue {
			inAttributeValue = true
		} else if r == '"' && inAttributeValue {
			inAttributeValue = false
		} else if !inElementTag || inAttributeValue {
			if relIdx := strings.Index(xml[i:], needle); relIdx > -1 {
				return i + relIdx
			}
		}

		i++
	}

	return -1
}

func main() {
	xmldoc := `<start serial="0xa5540012">
serial ...0012 has been issued
</start>`
	if FindInXMLString(xmldoc, "0012") > -1 {
		fmt.Println("Good: found substring '0012' as part of an element's content.")
	} else {
		fmt.Println("Bad: missed substring '0012' as part of an element's content.")
	}

	if FindInXMLString(xmldoc, "serial") > -1 {
		fmt.Println("Bad: found substring 'serial' as part of an attribute's name.")
	} else {
		fmt.Println("Good: ignored substring 'serial' as part of an attribute's name.")
	}

	if FindInXMLString(xmldoc, "start") > -1 {
		fmt.Println("Bad: found substring 'start' as part of an element's name.")
	} else {
		fmt.Println("Good: ignored substring 'serial' as part of an element's name.")
	}
}
