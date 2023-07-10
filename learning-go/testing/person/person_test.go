package person

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  27,
	}
	result := CreatePerson("Dennis", 27)

	// Utilize a custom comparer because we can't control DateAdded field
	comparer := cmp.Comparer(func(x Person, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
