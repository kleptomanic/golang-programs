// Sum numbers from 1 to 50.

package main

import "fmt"

func main() {
	fmt.Println("Sum of 1 to 50 natural number is : ", func(sum int) int {
		for i := 1; i <= 50; i++ {
			sum = sum + i
		}
		return sum

	}(0))
}
