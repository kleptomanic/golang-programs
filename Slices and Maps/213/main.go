// Check if a key exists in a map.

package main

import (
	"fmt"
)

func main() {
	userMap := map[string]string{
		"name":    "test",
		"age":     "32",
		"address": "Kathmandu",
	}
	val, _ := userMap["age"]
	if val != "" {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
