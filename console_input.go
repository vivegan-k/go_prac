package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered text", input)

	fmt.Print("Enter number:")
	numInput, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err == nil {
		fmt.Println("Entered number", number)
		fmt.Printf("Number type is %T", number)
	} else {
		fmt.Println(err)
	}

}
