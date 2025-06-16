// Swap two integer variables without a temp variable.

package main

import "fmt"

func main() {
	a, b := 2, 3
	fmt.Println("Values of variables when not swapped : \na ->", a, "\nb ->", b)

	// Swapping variables using temp value
	temp := a
	a = b
	b = temp

	fmt.Println("Values of variables when swapped using temp variable : \na ->", a, "\nb ->", b)

	// Swapping variables without using temp value
	a, b = 4, 5
	a, b = b, a
	fmt.Println("Values of variables when swapped using temp variable : \na ->", a, "\nb ->", b)

}
