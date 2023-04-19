package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"testing"
)

/* TestMain https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc
   We use this method to fire up the HTTP server to be used by all tests. */
func TestMain(m *testing.M) {
	go main() // startup HTTP server
	m.Run() // run tests
}

func TestSingleRequest(t *testing.T) {
	// arrange

	// act
	r, err := http.Get("http://localhost:2081/")
	if err != nil {
		t.Error(err)
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error(err)
		return
	}

	// assert
	num, err := strconv.Atoi(string(content))
	if err != nil {
		t.Error(fmt.Errorf("could not convert reponse to number: %v", err))
	}

	if num < 0 || num > 5000 {
		t.Error(fmt.Errorf("unexpected numeric value: %v", num))
	}
}

func TestMultiRequestsHammering(t *testing.T) {
	// arrange
	const n = 10
	ch := make(chan int, n) // don't block, so we can stress the HTTP server
	res := make([]bool, n)
	wg := sync.WaitGroup{}
	wg.Add(n)

	// act

	// prefetch first ID, so we now the offset (HTTP server is being used for other unit tests as well)
	offset, done := requestId(t)
	if done {
		return
	}
	offset++ // convert to slice index

	// hammer the HTTP server
	for i := 0; i < n; i++ {
		go func() {
			num, done := requestId(t)
			if done {
				return
			}

			ch <- num
			wg.Done()
		}()
	}

	wg.Wait()
	close(ch)

	// assert
	dups := make([]int, 0)
	for id := range ch {
		if res[id-offset] {
			dups = append(dups, id)
		}
		res[id-offset] = true // mark this id
	}
	if len(dups) > 0 {
		t.Error(fmt.Errorf("duplicate id's: %v", dups))
	}

	misses := make([]int, 0)
	for i := 0; i < n; i++ {
		if !res[i] {
			misses = append(misses, i+offset)
		}
	}
	if len(misses) > 0 {
		t.Error(fmt.Errorf("missing id's: %v", misses))
	}
}

func TestMultiRequestsKindly(t *testing.T) {
	// arrange
	const n = 10
	ch := make(chan int, n)
	res := make([]bool, n)

	// act

	// prefetch first ID, so we now the offset (HTTP server is being used for other unit tests as well)
	offset, done := requestId(t)
	if done {
		return
	}
	offset++ // convert to slice index

	// kindly ask the HTTP server, one request after the other
	for i := 0; i < n; i++ {
		func() {
			num, done := requestId(t)
			if done {
				return
			}

			ch <- num
		}()
	}

	close(ch)

	// assert
	dups := make([]int, 0)
	for id := range ch {
		if res[id-offset] {
			dups = append(dups, id)
		}
		res[id-offset] = true // mark this id
	}
	if len(dups) > 0 {
		t.Error(fmt.Errorf("duplicate id's: %v", dups))
	}

	misses := make([]int, 0)
	for i := 0; i < n; i++ {
		if !res[i] {
			misses = append(misses, i+offset)
		}
	}
	if len(misses) > 0 {
		t.Error(fmt.Errorf("missing id's: %v", misses))
	}
}

func requestId(t *testing.T) (int, bool) {
	r, err := http.Get("http://localhost:2081/")
	if err != nil {
		t.Error(err)
		return 0, true
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error(err)
		return 0, true
	}

	// assert
	num, err := strconv.Atoi(string(content))
	if err != nil {
		t.Error(fmt.Errorf("could not convert reponse to number: %v", err))
	}
	return num, false
}
