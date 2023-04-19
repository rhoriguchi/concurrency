package service

import (
	"unitTests3Mock/entity"
	"unitTests3Mock/storage"
)

func GetOld() entity.People {
	people := entity.People(storage.GetPersons())
	return people.FilterBy(func(p entity.Person) bool {
		return p.YearOfBirth < 1990
	})
}
