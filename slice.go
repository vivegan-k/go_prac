package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{1, 2, 7, 8, 3, 4, 5, 6}
	fmt.Println(arr)
	new_arr := append(arr, 7)
	fmt.Println(new_arr)

	arr_make := make([]int, 3)
	arr_make[0] = 1
	arr_make[1] = 2
	arr_make[2] = 3
	fmt.Println(arr_make)

	fmt.Println(arr_make[1:])

	sort.Ints(arr)
	fmt.Println(arr)

}
