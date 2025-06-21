// Sort a slice in ascending order.

package main

import (
	"fmt"
	"slices"
)

func main() {
	sli := []int{1, 100, 99, 55, 4, 20, 10}
	slices.Sort(sli)
	fmt.Println(sli)
}
