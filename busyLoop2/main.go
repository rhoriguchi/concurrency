package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO investigate why this program is using 100% CPU time

var wg sync.WaitGroup

func differential(a, b <-chan int) {
	sum := 0
	defer wg.Done()

	for {
		select {
		case v, ok := <-a:
			if !ok {
				return
			}

			sum += v

		case v, ok := <-b:
			if !ok {
				return
			}

			sum -= v

		default:
			// we wait for the sum to become even
			if sum != 0 && sum%2 == 0 {
				fmt.Printf("We struck gold! sum=%v\n", sum)
				return
			}
		}
	}
}

func main() {
	up := make(chan int)

	wg.Add(1)
	go differential(up, nil)

	for _, v := range []int{1, 2, 4, 6, 8, 10, 12, 13} {
		up <- v
		time.Sleep(1 * time.Second)
	}

	wg.Wait()
}
