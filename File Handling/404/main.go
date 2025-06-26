// Count lines in a text file.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	textFilePath := "testfile.txt"

	fileOpen, err := os.OpenFile(textFilePath, os.O_RDONLY, 0400)
	if err != nil {
		fmt.Println(err)
	}
	defer fileOpen.Close()
	buf := make([]byte, 2048)
	fileReadContentLength, err := fileOpen.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	stringFileContent := string(buf)
	fmt.Println(fileReadContentLength, " bytes of file content is read.")

	fmt.Println("Number of Lines: ", len(strings.Split(stringFileContent, "\n")))
}
