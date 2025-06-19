// Print a triangle of numbers (1 to 4)

package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
