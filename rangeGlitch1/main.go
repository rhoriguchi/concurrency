package main

import "fmt"

/*
	Run this code multiple times and compare the results.

	Side note: also compare running this with go run main.go and a compiled binary:
		1. i=0; while [ $i -lt 100 ]; do go run main.go; let i++; done
		2. go build; i=0; while [ $i -lt 100 ]; do ./rangeGlitch1; let i++; done
*/

// add env variable "GOEXPERIMENT=loopvar" to run configuration"
//
// export GOEXPERIMENT=loopvar; go build; i=0; while [ $i -lt 100 ]; do ./rangeGlitch1; let i++; done

func isEvenNumber(n int) bool {
	return n%2 == 0
}

func main() {
	ch := make(chan bool)

	jobs := []int{3, 887, 445, 1013, -45, 6}
	for _, v := range jobs {
		go func() {
			ch <- isEvenNumber(v)
		}()
	}

	var evenCount int
	var oddCount int
	for evenCount+oddCount < len(jobs) {
		if <-ch {
			evenCount++
		} else {
			oddCount++
		}
	}

	fmt.Printf("yes: %v, no: %v\n", evenCount, oddCount)
}
