package model

import (
	"reflect"
	"testing"
)

func TestAncestors(t *testing.T) {
	person := &Person{
		Name: "Aaron",
		Mother: &Person{
			Name: "Becky",
			Mother: &Person{
				Name: "Joanne",
			},
			Father: &Person{
				Name: "Donald",
			},
		},
		Father: &Person{
			Name: "Steve",
			Mother: &Person{
				Name: "Patricia",
			},
			Father: &Person{
				Name: "Charles",
			},
		},
	}

	expected := []string{"Aaron", "Becky", "Joanne", "Donald", "Steve", "Patricia", "Charles"}
	actual := person.Ancestors()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v but got: %v", expected, actual)
	}
}
