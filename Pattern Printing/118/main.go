// Print a 4x4 grid of decreasing numbers (16 to 1).

package main

import "fmt"

func main() {
	count := 16

	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Print(count, "\t")
			count--
		}
		fmt.Println()
	}
}
