// Combine two strings with a comma.

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"test", "test1"}
	fmt.Println(strings.Join(str, ","))
}
