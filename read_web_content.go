package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://api.coindesk.com/v1/bpi/currentprice.json"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Printf("Type of respone is %T", resp)
	fmt.Println(resp.Body)
	//resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)

	fmt.Println(content)

}
