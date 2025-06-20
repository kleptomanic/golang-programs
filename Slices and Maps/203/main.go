// Find the maximum in a slice.

package main

import (
	"fmt"
	"slices"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Max(list))
}
