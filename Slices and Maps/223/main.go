// Check if two slices are equal.

package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{5, 4, 3, 2, 1}
	// slices.Sort(s2)
	if slices.Compare(s1, s2) == 0 {
		fmt.Println("Same")
	} else {
		fmt.Println("Different")
	}
}
