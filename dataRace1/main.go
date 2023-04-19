package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
	1. Run this code with go run -race and point your browser to http://localhost:2081/
       It will return 0 and no indication of a data race.
	2. Now hit the browser's reload button a couple of times in a row VERY quickly
       Boom! A data race has been detected (watch the console output of this program).
       This is also demonstrated by the unit tests.

	Can you fix the data race?
*/
func main() {
	var nextId int

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		time.Sleep(time.Millisecond*200)
		w.Write([]byte(fmt.Sprintf("%v", nextId)))
		// revelation of a data race https://github.com/golang/go/blob/aee9a19c559da6fd258a8609556d89f6fad2a6d8/src/net/http/server.go#L3089
		nextId++
	})

	http.ListenAndServe("localhost:2081", http.DefaultServeMux)
}
