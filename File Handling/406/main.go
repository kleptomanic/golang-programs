// Check if a file exists.

package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	filePath := "file.txt"
	fileExists, err := os.Open(filePath)
	if pathError, ok := err.(*fs.PathError); ok {
		fmt.Println("File doesnot Exists", pathError.Path)
		// fmt.Printf("%T", err)
	} else {
		fmt.Println("File Exists", fileExists)
	}

}
