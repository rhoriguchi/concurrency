package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func produce(index int, data []int) {
	// simulate varying computational effort
	time.Sleep(time.Duration(2+rand.Intn(8)) * time.Millisecond)

	wg.Add(1)

	if index >= len(data) {
		log.Fatalln("Index %v out of range (0-%v)", index, len(data)-1)
	}

	data[index] = rand.Intn(100) + 1
	wg.Done()
}

func process(index int, data []int) {
	// simulate varying computational effort
	time.Sleep(time.Duration(2+rand.Intn(8)) * time.Millisecond)

	fmt.Printf("%02v: %v\n", index, data[index])
	if data[index] == 0 {
		fmt.Printf("Ooops! Missing product found at index %v\n", index)
	}

	wg.Done()
}

func main() {
	const size = 20
	d := make([]int, size)
	wg.Add(1)
	for i := 0; i < 20; i++ {
		go produce(i, d)
	}
	wg.Done()
	wg.Wait()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go process(i, d)
	}
	wg.Wait()
}
