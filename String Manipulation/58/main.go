// Count words in a string.

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Replace returns a copy of the string s with the first n non-overlapping-instances of old replaced by new."
	fmt.Println("Number of Words in a String are : ", len(strings.Split(str, " ")))

}
