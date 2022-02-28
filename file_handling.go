package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello Vivek"
	file, err := os.Create("./test.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("length of the string", length)
	defer file.Close()
	defer readFile("./test.txt")

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(file string) {
	data, err := ioutil.ReadFile(file)
	checkError(err)
	fmt.Println(string(data))
}
