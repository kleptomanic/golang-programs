// Add two integers and print their sum.

package main

import "fmt"

func main() {
	fmt.Println("Sum is : ", func(a, b int) int {
		return a + b
	}(1, 2))
}
