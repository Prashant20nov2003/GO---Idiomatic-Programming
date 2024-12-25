package main

import (
	"fmt"

	"github.com/Prashant20nov2003/ch10/sample_code/circular_dependency_example/person"
)

func main() {
	bob := person.Person{PetName: "Fluffy"}
	fmt.Println(bob.Pet())
}
