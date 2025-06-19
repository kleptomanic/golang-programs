// Find the index of a specific character.

package main

import "fmt"

func main() {
	str := "Find the index of a specific character."

	for i, s := range str {
		fmt.Println(i, " -> ", string(s))
	}
}
