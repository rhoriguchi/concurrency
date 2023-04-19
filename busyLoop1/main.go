package main

import "fmt"

// TODO fix the code, so it does not panic and does not loop forever

func merge(a, b <-chan int) chan int {
	out := make(chan int)

	go func() {
		for {
			select {
			case v := <-a:
				out <- v
			case v := <-b:
				out <- v
			}
		}
	}()

	return out
}

func launchProducer(data []int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range data {
			c <- v
		}
	}()

	return c
}

func main() {
	a := launchProducer([]int{9, 7, 5, 3, 1})
	b := launchProducer([]int{8, 6, 4, 2, 0})

	m := merge(a, b)

	for v := range m {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// source: GopherCon 2016: Francesc Campoy - Understanding nil https://www.youtube.com/watch?v=ynoY2xz-F8s&t=1460s
