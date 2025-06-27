// Copy one text file to another.

package main

import (
	"fmt"
	"os"
	"syscall"
)

const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	readOnly  int = syscall.O_RDONLY // open the file read-only.
	writeOnly int = syscall.O_WRONLY // open the file write-only.
	readWrite int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	append int = syscall.O_APPEND // append data to the file when writing.
	create int = syscall.O_CREAT  // create a new file if none exists.
)

type fileDetail struct {
	Name       string
	Flags      int
	Permission os.FileMode
	File       *os.File
}

func (f *fileDetail) assignFileValues(name string, flags int, permission os.FileMode) {
	f.Name = name
	f.Flags = flags
	f.Permission = permission
}

func (f *fileDetail) createFile() error {
	fileCreate, err := os.OpenFile(f.Name, f.Flags, f.Permission)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	f.File = fileCreate
	fmt.Println(fileCreate.Name(), "is successfully created.")
	return nil
}

func (f *fileDetail) writeContentsToFile(content string) error {
	if f.File == nil {
		return fmt.Errorf("%v", f.Name)
	}
	_, err := f.File.Write([]byte(content))
	fmt.Println(content)
	if err != nil {
		return fmt.Errorf("issue writing into a file %v", f.Name)
	}
	return nil
}

func (f *fileDetail) readDataFromFile() (string, error) {
	buf := make([]byte, 2048)
	if f.File == nil {
		return "", fmt.Errorf("empty File : %v", f.Name)
	}
	// Seek to the start of the file to ensure correct reading
	_, err := f.File.Seek(0, 0)
	if err != nil {
		return "", fmt.Errorf("failed to seek to start of file %s: %w", f.Name, err)
	}
	n, err := f.File.Read(buf)
	fmt.Println(string(buf), n)
	if err != nil {
		return "", fmt.Errorf("error reading from file : %v", err)
	}
	return string(buf), nil

}
func main() {
	fileOne := fileDetail{}
	fileTwo := fileDetail{}

	fileOne.assignFileValues("file1.txt", readWrite|create, 0600)
	fileTwo.assignFileValues("file2.txt", readWrite|create, 0600)

	fileOne.createFile()
	fileTwo.createFile()

	if err := fileOne.writeContentsToFile("Hello World"); err != nil {
		fmt.Println(err)
		return
	}
	data, err := fileOne.readDataFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := fileTwo.writeContentsToFile(data); err != nil {
		fmt.Println(err)
		return
	}
}
