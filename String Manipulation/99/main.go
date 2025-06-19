// Count lowercase letters in a string.

package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "Kaushik Raj Panta"
	count := 0
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if unicode.IsLower(runes[i]) {
			count++
		}
	}
	// Range must not be used for runes, But gives the expected result
	// for _, i := range runes {
	// 	if unicode.IsLower(i) {
	// 		count++
	// 	}
	// }
	fmt.Println("Count of Lower case letters : ", count)
}
