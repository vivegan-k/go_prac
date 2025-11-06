package main

import (
	"fmt"
)

const aConst = 10

func main() {
	fmt.Println("Hello World")
	// Variables
	var newstr string = "Test string" // Explicit typing

	fmt.Println(newstr)

	var newint int = 1 // Explicit typing

	fmt.Println(newint)

	i := 10 // Implict typing, this kind of assignment can be used only inside the function
	fmt.Println(i)
	fmt.Printf("Type is %T\n", i)

	var anotherStr = "Test string 2"
	fmt.Println(anotherStr)

	fmt.Println(aConst)
	anotherStr = "new one"
	fmt.Println(anotherStr)
}
