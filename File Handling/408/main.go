// Read a file and print lines in reverse.

package main

import (
	"fmt"
	"os"
)

type fileDetailsType struct {
	Name       string
	Permission os.FileMode
	Flags      int
	File       *os.File
}

func (f *fileDetailsType) initilizeFileDetails(name string, permission os.FileMode, flags int) {
	f.Name = name
	f.Permission = permission
	f.Flags = flags
}

func (f *fileDetailsType) openFile() {
	fileDetail, err := os.OpenFile(f.Name, f.Flags, f.Permission)
	if err != nil {
		fmt.Println("Issue openining the file.", err)
	}
	f.File = fileDetail
}

func (f *fileDetailsType) writeDataIntoFile(content string) error {
	_, err := f.File.Write([]byte(content))
	if err != nil {
		return fmt.Errorf("issue with writing file : %v", err)
	}
	fmt.Println("data Written into file successfully.")
	return nil
}

func (f *fileDetailsType) readDataFromFile() (string, error) {
	if f.File == nil {
		return "", fmt.Errorf("file is not defined : %v", f.Name)
	}
	_, err := f.File.Seek(0, 0)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	buf := make([]byte, 2048)
	dataIntread, err := f.File.Read(buf)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	fmt.Println(dataIntread)
	return string(buf), nil
}

func main() {
	fileName := "file1.txt"

	newFile := fileDetailsType{}
	newFile.initilizeFileDetails(fileName, 0600, os.O_RDWR|os.O_CREATE)
	newFile.openFile()
	if err := newFile.writeDataIntoFile("Hello World"); err != nil {
		fmt.Println(err)
	}
	returnedData, err := newFile.readDataFromFile()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Original String", returnedData)
	runes := []rune(returnedData)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	fmt.Println("Reversed String : ", string(runes))

}
