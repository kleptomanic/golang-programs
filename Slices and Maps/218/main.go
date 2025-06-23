// Calculate the average of a slice of floats.

package main

import "fmt"

func main() {
	list := []int{100, 45, 21, 33, 44}
	count := 0.0
	for i := 0; i < len(list); i++ {
		count = count + float64(list[i])
	}
	fmt.Println(count / float64(len(list)))
}
