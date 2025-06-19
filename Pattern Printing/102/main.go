// Print a right-angled triangle of stars (4 rows)

package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
