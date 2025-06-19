// Define a Person struct with name and age, then print it.

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := &Person{"Kushal", 30}
	fmt.Println(person.name, person.age)
}
