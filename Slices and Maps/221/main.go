// Find the minimum in a slice.

package main

import (
	"fmt"
	"slices"
)

func main() {
	sli := []int{-1, -10, 1, 55, 90}
	fmt.Println(slices.Min(sli))
}
