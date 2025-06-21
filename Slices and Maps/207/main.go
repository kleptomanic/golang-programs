// Check if a slice contains a number.

package main

import (
	"fmt"
	"strings"
)

func main() {
	sli := []string{"kaushik", "1", "kau7", "bibek"}

	if len(sli) < 1 {
		fmt.Println("Slice is empty.")
	}

	// Using contains any functions checks for number string
	for _, i := range sli {
		fmt.Println(strings.ContainsAny(i, "0123456789"))
	}

	// Using containsfunc
	fmt.Println()
	number := func(r rune) bool {
		return r == '1' || r == '2' || r == '3' || r == '4' || r == '5' || r == '6' || r == '7' || r == '8' || r == '9' || r == '0'
	}
	for _, i := range sli {
		fmt.Println(strings.ContainsFunc(i, number))
	}

}
