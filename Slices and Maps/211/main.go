// Create a map of names to ages and print it.

package main

import "fmt"

func main() {

	user := map[string]string{
		"name": "test",
		"age":  "32",
	}
	// fmt.Println(fmt.Println(user))
	for i, j := range user {
		fmt.Println(i, "->", j)
	}
}
