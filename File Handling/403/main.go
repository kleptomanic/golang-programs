// Write a string to a text file.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	str := "hello world"
	fileName, err := os.OpenFile("samplefile.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fileName.Close()

	dataWritten, err := fileName.Write([]byte(str))
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%v data is appended into file %v", dataWritten, fileName.Name())

}
