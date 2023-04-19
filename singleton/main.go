package main

import (
	"fmt"
	"singleton/oneOfAKind"
)

func main() {
	ch := make(chan int)

	go func(){
		ch <- oneOfAKind.GetId()
	}()
	go func(){
		ch <- oneOfAKind.GetId()
	}()

	fmt.Printf("id(1): %v\n", <-ch)
	fmt.Printf("id(2): %v\n", <-ch)
}
