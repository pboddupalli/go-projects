package backtracking

import "strings"

func solveNQueensVer1(n int) [][]string {
  var retval [][]string
  if n == 0 {
    return retval
  }
  rows := make([]int, n)
  cols := make([]bool, n)
  lookLeft := make([]bool, 2*n-1)
  lookRight := make([]bool, 2*n-1)
  solve(0, n, rows, cols, lookLeft, lookRight, &retval)
  return retval
}

func solve(row, n int, queenPosition []int, cols, lookLeft, lookRight []bool, retval *[][]string) {
  if row == n { // terminating condition
    entry := []string{}
    for row := 0; row < n; row++ {
      s := strings.Builder{}
      for col := 0; col < queenPosition[row]; col++ {
        s.WriteByte('.')
      }
      s.WriteByte('Q')
      for col := queenPosition[row]+1; col < n; col++ {
        s.WriteByte('.')
      }
      entry = append(entry, s.String())
    }
    *retval = append(*retval, entry)
    return
  }
  // for col := 1; col <= n/2 + 1; col++ {
  for col := 0; col < n; col++ {
    if !cols[col] && !lookLeft[n-row+col-1] && !lookRight[row+col] {
      cols[col], lookRight[row+col], lookLeft[n-row+col-1] = true, true, true
      queenPosition[row] = col
      solve(row+1, n, queenPosition, cols, lookLeft, lookRight, retval)
      cols[col], lookRight[row+col], lookLeft[n-row+col-1] = false, false, false
    }
  }
}
