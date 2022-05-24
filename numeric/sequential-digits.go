package numeric

import (
  "strconv"
)

func genNumbers(numDigits int, low, high int) []int {
  if numDigits < 1 || numDigits > 9 {
    return []int{}
  }
  result := make([]int, 0)
  for i := 1; i <= 10-numDigits; i++ {
    // k is the first digit. keep incrementing it for numDigits
    k := i
    tmp := make([]byte, 0)
    for j := 0; j < numDigits; j, k = j+1, k+1 {
      tmp = append(tmp, byte(k) + '0')
    }
    val, err := strconv.Atoi(string(tmp))
    if err == nil && val >= low && val <= high {
      result = append(result, val)
    }
  }
  return result
}

func getNumDigits(n int) int {
  cnt := 0
  for n > 0 {
    cnt++
    n = n / 10
  }
  return cnt
}

func sequentialDigits(low int, high int) []int {
  numLowDigits := getNumDigits(low)
  numHighDigits := getNumDigits(high)

  result := make([]int, 0)
  for i := numLowDigits; i <= numHighDigits; i++ {
    result = append(result, genNumbers(i, low, high)...)
  }
  return result
}
