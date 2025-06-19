// Count occurrences of "e" in a string.

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Count occurrences of 'e' in a string."
	fmt.Println(strings.Count(str, "e"))
}
