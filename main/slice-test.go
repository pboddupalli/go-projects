package main

import (
  "fmt"
  "sort"
)

func main() {
  slice := []int{1,2,3,4}
  /* secondSlice := slice
  slice = slice[len(slice):]
  fmt.Println(secondSlice)
  slice = append(slice, 3)
  fmt.Println(secondSlice) */

  sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j]} )
  fmt.Println(slice)
}
