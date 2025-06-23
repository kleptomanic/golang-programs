// Find the minimum in a slice.

package main

import (
	"fmt"
	"slices"
)

func main() {
	min := []int{65, 32, 78, 90, 43, 23, -1}
	fmt.Println(slices.Min(min))
}
