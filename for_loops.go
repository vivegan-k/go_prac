package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for i := range numbers {
		fmt.Println(numbers[i])
	}

	for index, val := range numbers {
		fmt.Println(index, val)
	}

	sum := 0
	for _, val := range numbers {
		sum += val
		if sum > 9 {
			goto theEnd
		}
	}
	fmt.Println(sum)
theEnd:
	fmt.Println("End of program", sum)

}
