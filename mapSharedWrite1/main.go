package main

/*
	TODO:
		1. Run this code multiple times until it aborts with panic "fatal error: concurrent map writes".
		2. Fix the issue.
*/

import (
	"fmt"
	"sync"
)

type Record struct {
	Id   int
	Name string
}

func main() {
	// create a map
	m := make(map[int]Record)

	// prepare and lock concurrent write access
	start := &sync.WaitGroup{}
	start.Add(1)

	done := &sync.WaitGroup{}
	done.Add(6)

	go WaitAddRecord(m, Record{1, "Anna Arglist"}, start, done)
	go WaitAddRecord(m, Record{2, "Bernd Brenner"}, start, done)
	go WaitAddRecord(m, Record{3, "Charlie Capman"}, start, done)
	go WaitAddRecord(m, Record{4, "Doris Dorn"}, start, done)
	go WaitAddRecord(m, Record{5, "Eberhard Einheim"}, start, done)
	go WaitAddRecord(m, Record{6, "Frederike Fernweh"}, start, done)

	// trigger write access
	start.Done()

	// make sure all go routines have completed
	done.Wait()

	// show me your hand
	fmt.Println(m)
}

func WaitAddRecord(m map[int]Record, r Record, wg *sync.WaitGroup, done *sync.WaitGroup) {
	// wait until we got the signal
	wg.Wait()

	// write access
	m[r.Id] = r

	// signal completion
	done.Done()
}
