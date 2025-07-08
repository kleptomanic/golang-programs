// Handle file not found error.

package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if perr, ok := err.(*fs.PathError); ok {
		log.Fatalln(perr.Unwrap().Error())
	}
	fmt.Println("File found.", file.Name())
}
