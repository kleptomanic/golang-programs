// Count true values in a slice of booleans.

package main

import "fmt"

func main() {
	boolValues := []bool{true, false, false, true, true}
	count := 0
	for i := 0; i < len(boolValues); i++ {
		if boolValues[i] {
			count += 1
		}
	}
	fmt.Print("Number of True values are : ", count)
}
