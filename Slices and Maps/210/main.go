// Merge two slices of integers.

package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 0}
	fmt.Println(slices.Concat(a, b))
}
