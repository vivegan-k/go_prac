package main

import (
	"fmt"
)

func main() {
	var result string

	num := 10

	if num%2 == 0 {
		result = "number is even"
	} else if num%2 != 0 {
		result = "odd"
	} else {
		result = "Invalid number"
	}
	fmt.Println(result)

	if num1 := 15; num1%2 == 0 {
		result = "number is even"
	} else if num1%2 != 0 {
		result = "odd"
	} else {
		result = "Invalid number"
	}
	fmt.Println(result)

}
