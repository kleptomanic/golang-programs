// Print a 4x4 square of stars.

package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
