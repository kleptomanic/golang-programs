// Create a slice of integers and print it.

package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	for _, i := range list {
		fmt.Println(i)
	}
}
