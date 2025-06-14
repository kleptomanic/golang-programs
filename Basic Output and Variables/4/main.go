// Declare a float64 and print it with two decimal places.

package main

import "fmt"

func main() {
	floatvar := 3.14151617

	decimal_var := fmt.Sprintf("%.2f", floatvar)
	fmt.Println(decimal_var)
}
