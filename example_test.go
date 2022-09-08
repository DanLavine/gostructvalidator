package gostructvalidator_test

import (
	"log"

	"github.com/DanLavine/gostructvalidator"
)

func Example() {
	type Person struct {
		Name     string
		Age      int
		Address  string
		Children []Person

		// This is used to just highlight how to validate a map object
		// map[ CITY ] Address
		PreviousAddresses map[string]string
	}

	person := Person{
		Name:    "Unamed Person",
		Age:     63,
		Address: "123 Over There Rd",
		Children: []Person{
			{Name: "child one", Age: 42},
			{Name: "child two", Address: "First person living on the moon"},
			{Name: "child three", Age: 28, Children: []Person{
				{Name: "grand child 1", Age: 3},
			}},
		},
	}

	structValidator := gostructvalidator.New()
	if err := structValidator.Validate(person); err != nil {
		log.Fatal(err)
	}
}
