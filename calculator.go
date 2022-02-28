package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number1: ")
	input1, _ := reader.ReadString('\n')
	number1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err1 != nil {
		panic("Entered input is not integer")
	}
	fmt.Print("Enter the number2: ")
	input2, _ := reader.ReadString('\n')
	number2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err2 != nil {
		panic("Entered input is not integer")
	}
	fmt.Println(math.Round(number1 + number2))
}
