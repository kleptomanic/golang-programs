// Print all keys in a map.

package main

import (
	"fmt"
	"maps"
)

func main() {
	userData := map[string]string{
		"name": "test",
		"age":  "32",
	}
	for i, j := range userData {
		fmt.Println(i, j)
	}
	fmt.Println(maps.Keys(userData))
}
