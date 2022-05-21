package main

import (
  "fmt"
  "strings"
)

func main() {
  s := "abc"
  s = strings.Repeat(s, 3)
  fmt.Println(s)
  s = "abc"
  s = strings.Replace(s, "", "d", -1)
  fmt.Println(s)
}
