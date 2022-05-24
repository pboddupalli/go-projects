package paranthesis

import "fmt"

func longestValidParanthesis(s string) int {
  result := 0
  open, close1 := 0, 0
  stack := make([]rune, 0)

  for _, c := range s {
    if byte(c) == '(' {
      open++
    } else {
      open--
    }
    if open < 0 {
      // empty the stack
      result = max(result, len(stack))
      stack = stack[len(stack):]
      open = 0
      // ignore close1 bracket
    } else {
      stack = append(stack, c)
    }
  }
  fmt.Println(len(stack), result)

  // process stack backwards
  sum := 0
  for len(stack) != 0 {
    c := stack[len(stack)-1]
    stack = stack[len(stack)-1:] // pop
    if byte(c) == ')' {
    close1++
    } else {
    close1--
    }
    if close1 < 0 {
      result = max(result, sum)
      sum = 0
    } else {
      sum++
    }
  }
  result = max(result, sum)
  return result
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}


