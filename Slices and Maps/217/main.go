// Count character frequencies in a string using a map.

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is a random character sentences"
	countMap := make(map[string]int)

	for _, j := range str {
		countMap[string(j)] = strings.Count(str, string(j))
	}
	fmt.Println("Char", "count")
	for i, j := range countMap {
		fmt.Println(i, "=>\t", j)
	}

}
