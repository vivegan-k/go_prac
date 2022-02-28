package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)
	var result string

	switch dow {
	case 1:
		result = "one"
	case 2:
		result = "two"
	default:
		result = "other numbers"
	}
	fmt.Println(result)
}
