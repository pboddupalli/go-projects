package memoization

func longestLine(mat [][]int) int {

  result := 0

  m, n := len(mat), len(mat[0])
  prev := make([][4]int, n + 2)
  current := make([][4]int, n + 2)

  for i := 0; i < m; i++ {
    prev, current = current, prev
    for j := 1; j <= n; j++ {
      if mat[i][j-1] == 0 {
        // reset them to zero...
        for k := 0; k < 4; k++ { current[j][k] = 0 }
      } else {
        current[j][0] = prev[j][0] + 1
        current[j][1] = current[j-1][1] + 1
        current[j][2] = prev[j-1][2] + 1
        current[j][3] = prev[j+1][3] + 1
      }
      retval := max(current[j])
      if retval > result {
        result = retval
      }
    }
  }
  return result
}

func max(nums [4]int) int {
  result := 0
  for _, val := range nums {
    if val > result {
      result = val
    }
  }
  return result
}
