// Parse a string to int and handle errors.

package main

import (
	"log"
	"strconv"
)

func main() {
	str := "test"

	if _, err := strconv.Atoi(str); err != nil {
		log.Fatalf("%v", err)
	}
}
