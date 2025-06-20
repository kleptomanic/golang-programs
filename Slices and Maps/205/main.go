// Remove an element by index from a slice.

package main

import (
	"fmt"
	"slices"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	slices.Delete(list, 3, 4)
	// removing the trailing 0 in slice
	fmt.Println(list[:len(list)-1])
}
