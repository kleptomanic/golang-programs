// Replace "a" with "@" in a string.

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Replace a with @ in a string."
	fmt.Println(strings.ReplaceAll(str, "a", "@"))

}
