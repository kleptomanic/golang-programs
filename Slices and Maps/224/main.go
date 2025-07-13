// Double each element in a slice.

package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 5, 6}
	for i := range sli {
		sli[i] *= 2
	}
	fmt.Println(sli)
}
