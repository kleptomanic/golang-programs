// Print an inverted right-angled triangle (4 rows).

package main

import "fmt"

func main() {
	for i := 4; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
