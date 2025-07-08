// Divide two numbers and handle division by zero.
package main

import (
	"fmt"
	"os"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, errors.New("cannot be divisible by 0")
		return 0, fmt.Errorf("cannot be divisible by %v", b)
	}
	return a / b, nil
}

func main() {
	a, b := 2, 0
	result, err := divide(a, b)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(result)
}
