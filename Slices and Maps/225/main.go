// Find the index of an element in a slice.

package main

import (
	"fmt"
	"slices"
)

func main() {
	in := []int{1, 5, 6, 7, 899, 0, 44, 2}

	fmt.Println(slices.Index(in, 0))
}
