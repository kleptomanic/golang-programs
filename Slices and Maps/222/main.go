// Sort a slice of strings alphabetically.

package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	str := []string{"kush", "aanita", "anita", "nabin", "lopchan"}
	slices.SortFunc(str, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(str)
}
