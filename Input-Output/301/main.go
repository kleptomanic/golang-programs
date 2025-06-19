// Read a number and print its square.

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the Number : ")
	fmt.Scanf("%d", &num)
	fmt.Printf("Square of %d is %d.\n", num, num*num)
}
