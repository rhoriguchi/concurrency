package main

import (
	"fmt"
	"sync"
	"time"
)

/*
   Challenge: Implement concurrency in pingAll() to complete all jobs as fast as possible.
   Rule:      Only change main.go, do not change ANY OTHER file.
   Definition of done: ALL unit tests in this project show green, ping/*_test.go included.
*/

import (
	"concurrent2/ping"
)

func pingAll(jobs chan ping.Host) (int, int) {
	reachableCount := make(chan bool)

	var wg sync.WaitGroup

	for host := range jobs {
		wg.Add(1)

		go func(h ping.Host) {
			defer wg.Done()

			h.Ping()
			reachableCount <- h.Reachable != nil && *h.Reachable
		}(host)
	}

	go func() {
		wg.Wait()
		close(reachableCount)
	}()

	totalReachable := 0
	totalJobs := 0
	for r := range reachableCount {
		totalJobs++
		if r {
			totalReachable++
		}
	}

	return totalReachable, totalJobs
}

const hostsFilename = "hosts.csv.bz2"

func main() {
	startProg := time.Now()
	jobCh := ping.GetJobs(hostsFilename)

	startPing := time.Now()
	fmt.Printf("loaded hosts from %v\n", hostsFilename)

	reachableCount, hostCount := pingAll(jobCh)
	stopPing := time.Now()

	fmt.Printf("TIMING %v to parse %v hosts\n", startPing.Sub(startProg), hostCount)
	fmt.Printf("TIMING %v to ping %v hosts\n", stopPing.Sub(startPing), hostCount)
	fmt.Printf("RESULT %v/%v (%.6f%%) hosts reachable\n", reachableCount, hostCount, float64(reachableCount)/float64(hostCount)*100)
}
