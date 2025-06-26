// Count words in a text file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "testfile.txt"
	checkFileExistance, err := os.OpenFile(fileName, os.O_RDONLY, 0400)
	if err != nil {
		fmt.Println(err)
	}
	wordCount := 0
	scanner := bufio.NewScanner(checkFileExistance)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Skip empty lines
			words := strings.Fields(line)
			wordCount += len(words)
		}
	}
	fmt.Println(wordCount)
}
