// Return an error for negative input

package main

import "fmt"

func isNegative(num int) (int, error) {
	if num < 0 {
		return 0, fmt.Errorf("negative number detected : %v", num)
	}
	return num, nil
}

func main() {
	numbers := []int{1, -1, 5, 9, -3}
	for i, num := range numbers {
		result, err := isNegative(num)
		if err != nil {
			fmt.Printf("Error at index %v: %v\n", i, err)
			continue
		}
		fmt.Printf("Positive Number found at index %v: %v\n", i, result)
	}
}
