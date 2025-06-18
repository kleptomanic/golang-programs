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

	// f := func(r rune) bool {
	// 	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	// }
	// fmt.Println(strings.ContainsFunc("hello", f))
	// fmt.Println(strings.ContainsFunc("rhythms", f))
}
