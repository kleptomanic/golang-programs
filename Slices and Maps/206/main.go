// Reverse a slice of integers.

package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6}

	// Using builtin function
	// slices.Reverse(list)

	// Old style
	left := 0
	right := len(list) - 1
	for left < right {
		temp := list[left]
		list[left] = list[right]
		list[right] = temp
		left++
		right--
	}
	fmt.Println(list)
}
