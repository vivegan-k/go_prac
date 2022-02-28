package main

import (
	"fmt"
)

func main() {
	emp := Employee{"vivek", 30}
	fmt.Println(emp)
	emp.age = 31
	fmt.Println(emp)

}

type Employee struct {
	name string
	age  int
}
