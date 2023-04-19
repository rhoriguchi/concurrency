package entity

import (
	"fmt"
	"sort"
)

type Person struct {
	Id          int
	Name        string
	YearOfBirth int
}

func (p Person) String() string {
	return fmt.Sprintf("#%v %v *%v", p.Id, p.Name, p.YearOfBirth)
}

type People []Person
type byId []Person

func (people byId) Len() int {
	return len(byId{})
}

func (people byId) Swap(i, j int) {
	people[i], people[j] = people[j], people[i]
}

func (people byId) Less(i, j int) bool {
	return people[i].Id < people[j].Id
}

func SortById(people []Person) {
	sort.Sort(byId(people))
}

type byIdName struct {
	byId
}

func (people byIdName) Less(i, j int) bool {
	if people.byId[i].Id == people.byId[j].Id {
		return people.byId[i].Name < people.byId[j].Name
	}
	return people.byId.Less(i, j)
}

func SortByIdName(people []Person) {
	sort.Sort(byIdName{byId(people)})
}
