// Add a key-value pair to a map.

package main

import "fmt"

func main() {
	userMap := make(map[string]string)
	userMap["name"] = "test"
	userMap["age"] = "32"

	for i, j := range userMap {
		fmt.Println(i, j)
	}
}
