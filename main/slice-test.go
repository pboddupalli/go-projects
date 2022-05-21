package main

import "fmt"

func main() {
  slice := []int{1,2,3,4}
  secondSlice := slice
  slice = slice[len(slice):]
  fmt.Println(secondSlice)
  slice = append(slice, 3)
  fmt.Println(secondSlice)
}
