package main

import (
	"fmt"
)

func main() {
	maps := make(map[string]string)
	maps["name"] = "vivek"
	maps["age"] = "thirty"
	maps["job"] = "QE"
	fmt.Println(maps)

	fmt.Println(maps["job"])
	delete(maps, "job")
	fmt.Println(maps)

	for k, v := range maps {
		fmt.Println(k, v)
	}

	var keys = make([]string, len(maps))
	i := 0
	for k := range maps {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

}
