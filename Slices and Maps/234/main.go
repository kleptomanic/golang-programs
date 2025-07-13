// Check if a slice is sorted.

package main

import (
	"fmt"
	"slices"
)

func main() {
	sli := []int{45, 1, 4, 5, 90, 8, 41}

	fmt.Println(slices.IsSorted(sli))
}
