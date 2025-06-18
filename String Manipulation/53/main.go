// Count vowels in a string.

package main

import (
	"fmt"
	"strings"
)

func main() {
	vowels := "aeiouAEIOU"

	str := "kUshal"
	count := 0
	for _, char := range str {
		if strings.Contains(vowels, string(char)) {
			count++
		}
	}
	fmt.Println("Number of Vowels count :", count)

}
