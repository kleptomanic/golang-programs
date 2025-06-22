// Delete a key from a map.

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
	maps.DeleteFunc(userData, func(k string, v string) bool {
		return k == "name"
	})
	fmt.Println(userData)
}
