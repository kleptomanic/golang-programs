// Read and print a text file’s contents.

// Write a string to a text file.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileRead, err := os.OpenFile("testfile.txt", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer fileRead.Close()
	buf := make([]byte, 2048)

	// Buffer will include the data read from the file.
	fileReadData, err := fileRead.Read(buf)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(fileReadData, "bytes are read from file ", fileRead.Name())
	fmt.Println(string(buf))

}
