package paranthesis

import "fmt"

func longestValidParenthesesVer3(s string) int {
  stack := make([]int, 0)
  stack = append(stack, -1)
  result := 0
  for idx, c := range s {
    if c == '(' {
      stack = append(stack, idx)
      continue
    }
    // c == ‘)’. either there is a matching ‘(‘ or we need to make this index as the boundary
    top := stack[len(stack)-1]
    if top == -1 {
      stack = append(stack, idx)
    } else if s[top] == ')' {
      // overwrite it
      stack[len(stack)-1] = idx
    } else {
      stack = stack[0:len(stack)-1]
      if result < idx - stack[len(stack)-1] {
        result = idx - stack[len(stack)-1]
      }
    }
  }
  return result
}

func longestValidParenthesesVer2(s string) int {
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

  fmt.Println(len(stack))
  // process stack backwards
  sum := 0
  for len(stack) > 0 {
    c := stack[len(stack)-1]
    stack = stack[:len(stack)-1] // pop
    if byte(c) == ')' {
      close1++
    } else {
      close1--
    }
    if close1 < 0 {
      result = max(result, sum)
      sum = 0
      close1 = 0
    } else {
      sum++
    }
    fmt.Println("close1: sum:", close1, sum)
  }
  fmt.Println(sum)
  result = max(result, sum)
  return result
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}


func longestValidParentheses(s string) int {
  stack := make([]int, 0) // interesting...stacking index
  stack = append(stack, -1)
  maxLen := 0
  for i, c := range s {
    if c == '(' {
      stack = append(stack, i)
      continue
    }
    if len(stack) == 1 {
      // there is nothing we can do. push the index
      stack = append(stack, i)
    } else if s[stack[len(stack)-1]] == ')' {
      // top-of-stack is ')'
      stack[len(stack)-1] = i
    } else {
      // top of stack is '('
      stack = stack[:len(stack)-1] // ...pop it
      len := i - stack[len(stack)-1]
      if len > maxLen {
        maxLen = len
      }
    }
  }
  return maxLen
}
