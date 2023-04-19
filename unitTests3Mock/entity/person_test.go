package entity

import "testing"

func Test_SortById(t *testing.T) {
	// arrange
	people := []Person{
		{10, "Anna Caran", 2002},
		{4, "John Selvik", 1982},
		{12, "Abel Bamert", 1999},
	}

	// act
	SortById(people)

	// assert
	lastId := -1
	for i, p := range people {
		if p.Id < lastId {
			t.Errorf("not sorted ascending by id (index %v)", i)
		}
		lastId = p.Id
	}
}

func Test_SortByIdName(t *testing.T) {
	// arrange
	people := []Person{
		{10, "Abel Bamert", 1999},
		{4, "John Selvik", 1982},
		{10, "Anna Caran", 2002},
	}

	// act
	SortByIdName(people)

	// assert
	lastName := ""
	lastId := -1
	for i, p := range people {
		if p.Id < lastId {
			t.Errorf("not sorted ascending by id (index %v)", i)
		}
		
		if lastId == p.Id && lastName > p.Name {
			t.Errorf("not sorted ascending by name (index %v)", i)
		}

		lastId = p.Id
		lastName = p.Name
	}
}
