// Print the first 10 even numbers.

package main

import "fmt"

func main() {
	count := 10
	for i := 1; count > 0; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			count--
		}
	}
}
