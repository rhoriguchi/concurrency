package main

import (
	"fmt"
	"time"
)

const WORKERS = 20

type Result struct {
	Index   int
	Outcome bool
}

func processSingle(i, j int) bool {
	time.Sleep(16 * time.Microsecond)
	return i < j
}

func processConcurrently(all []int) []bool {
	in := make(chan []int)
	out := make(chan Result)

	// spin up workers
	for k := 0; k < WORKERS; k++ {
		go func() {
			for pair := range in {
				out <- Result{
					Index:   pair[0],
					Outcome: processSingle(pair[1], pair[2]),
				}
			}
		}()
	}

	// feed workers with tasks
	go func() {
		for i := 0; i < len(all); i += 2 {
			in <- []int{i/2, all[i], all[i+1]}
		}
	}()

	// collect outcome
	ret := make([]bool, len(all)/2)
	for r := range out {
		ret[r.Index] = r.Outcome
	}

	return ret
}

func main() {
	in := []int{1, 2, 4, 3, 22, -8, -55, 0}
	out := processConcurrently(in)
	fmt.Println(out)
}
