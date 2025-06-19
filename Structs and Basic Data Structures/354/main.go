// Define a Book struct with title and author.

package main

import "fmt"

type Book struct {
	title  string
	author string
}

func main() {
	book := []Book{
		{
			title:  "Learning Go",
			author: "Ram Hari Dangol",
		},
		{
			title:  "Practical Go",
			author: "Sita Dangol",
		},
	}
	for _, i := range book {
		fmt.Println(i.title, " -> ", i.author)
	}
}
