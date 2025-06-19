// Print a 4x4 grid of alternating 1s and 0s.

package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Print(j % 2)
		}
		fmt.Println()
	}
}
