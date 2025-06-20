// Sum elements in a slice.

package main

import "fmt"

func main() {
	count := 0
	list := []int{1, 2, 3, 4, 5}
	for _, i := range list {
		count += i
	}
	fmt.Println(count)
}
