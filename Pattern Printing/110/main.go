// Print a 4x4 grid of increasing numbers (1 to 16).

package main

import "fmt"

func main() {
	count := 1
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Print(count, "\t")
			count++
		}
		fmt.Println()
	}
}
