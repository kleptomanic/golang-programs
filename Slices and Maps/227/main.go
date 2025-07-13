// Remove duplicates from a slice.

package main

import (
	"fmt"
	"slices"
)

func main() {
	sli := []string{"test", "test", "test123", "test456", "anskdas"}

	fmt.Println(slices.Compact(sli))
}
