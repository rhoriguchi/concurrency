package service

import (
	"fmt"
	"unitTests3Mock/entity"
	"unitTests3Mock/storage"
)

func ShowStorage() {
	stg := storage.GetPersons()

	for i, p := range stg {
		fmt.Printf("-%v- %v\n", i, p)
	}
}

func Show(p entity.People) {
	for i, p2 := range p {
		fmt.Printf("-%v- %v\n", i, p2)
	}
}
