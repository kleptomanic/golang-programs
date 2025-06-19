// Create a slice of Person structs and print names.

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := []Person{
		{name: "ram", age: 30},
		{name: "hari", age: 40},
	}

	for _, i := range person {
		fmt.Println(i.name)
	}
}
