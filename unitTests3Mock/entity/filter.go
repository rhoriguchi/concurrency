package entity

func (p People) FilterBy(filter func(p Person) bool) People {
	var out People
	for _, person := range p {
		if filter(person) {
			out = append(out, person)
		}
	}

	return out
}
