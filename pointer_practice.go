package main

import (
	"fmt"
)

func main() {
	var newInt = 9
	pointernewInt := &newInt
	fmt.Println(*pointernewInt)
	fmt.Println(newInt)

	*pointernewInt = 10
	fmt.Println(*pointernewInt)
	fmt.Println(newInt)

}
