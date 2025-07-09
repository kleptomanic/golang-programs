// Handle negative input for square root

package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrtNum(num int) (float64, error) {
	if num < 0 {
		return 0, errors.New("negative number passed")
	}
	return math.Sqrt(float64(num)), nil
}

func main() {
	num := []int{1, 9, 45, -9}

	for _, n := range num {
		sqrtResult, err := sqrtNum(n)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Square root of %v is %v. \n", n, sqrtResult)
	}
}
