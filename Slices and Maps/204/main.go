// Append an element to a slice.

package main

import "fmt"

func main() {
	list := make([]int, 5)
	for i := 0; i < 5; i++ {
		list = append(list, i)
	}
	fmt.Println(list)
}
