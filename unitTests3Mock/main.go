package main

import (
	"fmt"
	"unitTests3Mock/entity"
	"unitTests3Mock/service"
	"unitTests3Mock/storage"
)

func main() {
	people := []entity.Person{
		{1, "Petra Power", 1997},
		{2, "Peter Power", 2001},
		{4, "Paul Power", 1961},
		{5, "Prada Power", 1967},
	}

	err := storage.Store(people)
	if err != nil {
		panic(fmt.Errorf("storing a family of 4 failed: %v", err))
	}

	fmt.Println("All people:")
	service.ShowStorage()

	filtered := entity.People(people).FilterBy(func(p entity.Person) bool {
		return p.YearOfBirth > 1990
	})

	fmt.Println("Filtered (born after 1990):")
	service.Show(filtered)

	old := service.GetOld()
	fmt.Printf("Found %v old people\n", len(old))
}
