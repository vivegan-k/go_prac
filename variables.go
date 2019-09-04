package main

import (
  "fmt"
  "strings"
)

func main(){
  var m1 int
  fmt.Println(m1) //By default int value is 0
  m1 = 2
  fmt.Println(m1)
  var ( //To create multiple variables
    a = 2
    b = 3
  )
  fmt.Println(a + b)
  m2 := 5
  fmt.Println(m1 + m2)


  //String values are mutable in go
  var s1 string
  fmt.Println(s1)

  s2 := "string2"
  fmt.Println(s2)

  s3 := "str"
  fmt.Println(strings.Contains(s2,s3))

  fmt.Println(strings.ReplaceAll(s2,"s","no"))
  fmt.Println(strings.Split("my name", " "))

  //Array declarations
  var arr1 []int
  fmt.Println(arr1)

  arr2 := []int{1,2,3,4}
  fmt.Println(arr2)

  arr3 := append(arr2, 4, 5, 6)
  fmt.Println(arr3)
  fmt.Println(arr3[0])
}
