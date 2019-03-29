package model

type Person struct {
	Name   string  `json:"name"`
	Mother *Person `json:"mother"`
	Father *Person `json:"father"`
}

func (person *Person) Ancestors() []string {
	var names []string
	if person != nil {
		names = append(names, person.Name)
		names = append(names, person.Mother.Ancestors()...)
		names = append(names, person.Father.Ancestors()...)
	}
	return names
}
