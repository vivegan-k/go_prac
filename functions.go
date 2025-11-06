package main

import (
	"fmt"
)

func main() {
	printSomething()
	sum := addValues(3, 5)
	fmt.Println(sum)
	total := multiAddition(1, 2, 3)
	fmt.Println(total)
	sum1, length := multiAdditionwithlength(1, 2, 3, 4)
	fmt.Println(sum1, length)
	emp := Employee{"vivk", 30, "QE"}
	emp.getJob()

}

func printSomething() {
	fmt.Println("Printing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func multiAddition(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func multiAdditionwithlength(values ...int) (int, int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, len(values)
}

// Example for methods

type Employee struct {
	name string
	age  int
	work string
}

func (e Employee) getJob() {
	fmt.Println("Working as ", e.work)
}
