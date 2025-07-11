// Check for map key existence.

package main

import "fmt"

func main() {
	testmap := map[string]string{
		"1": "name",
		"2": "address",
	}
	value, ok := testmap["3"]
	if !ok {
		fmt.Println("No Any data found in map")
	} else {
		fmt.Println(value)
	}
}
