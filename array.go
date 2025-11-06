package main

import (
	"fmt"
)

func main() {
	var arr [3]string
	arr[0] = "one"
	arr[1] = "two"
	arr[2] = "three"
	fmt.Println(arr, len(arr))

	num_arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(num_arr, len(num_arr))
}
