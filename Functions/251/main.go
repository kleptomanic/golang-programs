// Write a function to cube a number.

package main

import "fmt"

func cube(x int) int {
	return x * x * x
}

func main() {
	fmt.Println(cube(3))
}
