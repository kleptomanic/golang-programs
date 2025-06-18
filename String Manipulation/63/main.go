// Check if a string contains "Go".

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Gopher"
	str2 := "testGao"

	fmt.Println(strings.Contains(str1, "Go"))
	fmt.Println(strings.Contains(str2, "Go"))
}
