// Check if a string is empty.

package main

import (
	"fmt"
)

func main() {
	sample_string := ""
	if len(sample_string) == 0 {
		fmt.Println("Empty String")
	} else {
		fmt.Println("Non Empty String")
	}
}
