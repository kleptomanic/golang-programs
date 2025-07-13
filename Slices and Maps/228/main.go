// Concatenate a slice of strings.

package main

import (
	"fmt"
	"slices"
)

func main() {
	sli1 := []string{"a", "b", "c"}
	sli2 := []string{"q", "w", "r"}
	fmt.Println(slices.Concat(sli1, sli2))
}
