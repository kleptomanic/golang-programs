// Calculate the area of a rectangle.

package main

import "fmt"

func main() {
	fmt.Println("Area of rectangle : ", func(l, b int) int {
		return l * b
	}(3, 2))
}
