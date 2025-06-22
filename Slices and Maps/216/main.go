// Print all values in a map.

package main

import "fmt"

func main() {
	userData := map[string]string{
		"name": "test",
		"age":  "32",
	}
	for _, j := range userData {
		fmt.Println(j)
	}
}
