// Replace spaces with hyphens in a string.

package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new."
	fmt.Println(strings.ReplaceAll(name, " ", "-"))
}
