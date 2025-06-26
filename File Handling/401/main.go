// Write a string to a text file.

package main

import (
	"fmt"
	"os"
)

func main() {
	str := "test"
	fileName, err := os.OpenFile("samplefile.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fileName.Close()
	writeToFile, err := fileName.WriteString(str)
	if err != nil {
		fmt.Println("Issue with writing into file", err)
	}
	fmt.Println(writeToFile, "Bytes are written into file")

}
