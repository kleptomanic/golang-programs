// Check if two strings are equal (case-sensitive).

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Case Sensitive
	// If 0 then both are equal else Not Equal
	fmt.Println(strings.Compare("Nepal", "nepal"))

	// Case Insensitive compare
	// If True then, strings are equal else not equal
	fmt.Println(strings.EqualFold("Nepal", "nepal"))
}
