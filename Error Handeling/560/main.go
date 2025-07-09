// Check for valid slice index.

package main

import (
	"errors"
	"fmt"
)

func getElement(slice []string, index int) (string, error) {
	if index < 0 || index >= len(slice) {
		return "", errors.New("index out of bounds")
	}
	return slice[index], nil
}

func main() {
	slice := []string{"apple", "banana", "cherry"}
	indicesToCheck := []int{1, 3, -1}

	for _, index := range indicesToCheck {
		value, err := getElement(slice, index)
		if err != nil {
			fmt.Printf("Error accessing index %d: %v\n", index, err)
			continue
		}
		fmt.Printf("Value at index %d: %s\n", index, value)
	}
}
