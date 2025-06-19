// Update a Person’s age in a struct.

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := &Person{"Ram", 30}

	person.age = 45

	fmt.Println(person.age)
}
