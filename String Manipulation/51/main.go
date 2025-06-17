// Check if a string is empty.

package main

import (
	"fmt"
	"strings"
)

func main() {
	sample_string := " asdas"
	if strings.TrimSpace(sample_string) == "" {
		fmt.Println("Empty string is found.")
	} else {
		fmt.Println("Non Empty string is found ->", sample_string)
	}
}
