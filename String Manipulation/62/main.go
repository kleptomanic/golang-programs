// Print each character of a string on a new line.

package main

import "fmt"

func main() {
	str := "Print each character of a string on a new line."
	for _, i := range str {
		fmt.Println(string(i))
	}
}
