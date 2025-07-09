// Check if a string is a valid number

package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := []string{"123", "test123"}
	for _, s := range str {
		convertedInt, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Not the valid Number is found : ", s)
			continue
		}
		fmt.Println("Valid Number found : ", convertedInt)
	}
}
