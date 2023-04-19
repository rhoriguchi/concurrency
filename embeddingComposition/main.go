package main

import "fmt"

type I1 interface {
	GetElements() []int
	Put(v int)
}

type I2 interface {
	I1
	GetElementsFloat() []float32
}

type A struct {
	b int
}

type B struct {
	A
	c int
	b float32
}

func (a *A) Put(v int) {
	a.b = v
}

func (a A) GetElements() []int {
	return []int{a.b}
}

func (b B) GetElements() []int {
	return []int{b.c}
}

func (b B) GetElementsFloat() []float32 {
	return []float32{b.b}
}

func main() {

	a := A{b: 1}
	b := B{A{b: 7}, 2, 3.8}
	//b.A.b = 7

	fmt.Println(a)
	fmt.Println(b)

	b.Put(11)

	fmt.Println()
	fmt.Println(a.GetElements())

	// similar to function override: most specific method wins
	fmt.Println(b.GetElements(), b.A.GetElements())
	fmt.Println(b.GetElementsFloat())
}
