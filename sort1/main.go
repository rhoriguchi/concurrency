package main

/*
	TODO sort ascending by id
	TODO reverse sort using sort.Reverse()
 */

type Person struct {
	id int
	name string
	yearofbirth uint16
}

func main() {
	people := []Person{
		{9012, "John Connor", 1993},
		{71, "Sarah Connor", 1951},
		{431, "Elton Connor", 2002},
	}

	lastId := -1
	for _, p := range people {
		if p.id < lastId {
			panic("not sorted by id")
		}
		lastId = p.id
	}
}
