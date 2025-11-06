package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println(n)

	t := time.Date(2018, 8, 11, 16, 16, 1, 1, time.Local)
	fmt.Println("Iniyan was born on", t)
}
