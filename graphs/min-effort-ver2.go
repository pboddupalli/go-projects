package graphs

func minimumEffortPath(board [][]int) int {
  m,n := len(board), len(board[0])

  maxEffort := determineMaxEffort(board)
  if maxEffort == 0 {
    return 0
  }

  i, j := 0, maxEffort+1
  result := maxEffort
  for i < j {
    k := i + (j - i) / 2 // this is the binary value we want to travers the board on
    seen := make([][]int, m)
    for i := 0; i < m; i++ {
      seen[i] = make([]int, n)
    }
    canMeet := traverse(board, seen, 0, 0, k)
    if canMeet == 1 {
      result = k
      // we will try with even a smaller value
      j = k
    } else {
      i = k + 1
    }
  }
  return result
}

func traverse(board [][]int, seen [][]int, x, y int, k int) int {

  directions := [][]int{{-1,0}, {1,0}, {0,1},{0,-1}}
  m, n := len(seen), len(seen[0])
  if x == m-1 && y == n-1 {
    return 1
  }
  if seen[x][y] != 0 {
    return seen[x][y]
  }
  seen[x][y] = -1
  for _, dir := range directions {
    x1, y1 := x + dir[0], y + dir[1]
    if x1 < 0 || x1 >= m || y1 < 0 || y1 >= n {
      continue
    }
    if abs(board[x][y] - board[x1][y1]) > k {
      continue
    }
    childResult := traverse(board, seen, x1, y1, k)
    if seen[x][y] == -1 && childResult == 1 {
      seen[x][y] = 1
      break
    }
  }
  return seen[x][y]
}

func determineMaxEffort(board [][]int) int {
  result := 0
  m,n := len(board), len(board[0])

  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      i1, j1 := i+1, j
      if i1 < m && j1 < n {
        result = max(result, abs(board[i][j] - board[i1][j1]))
      }
      i2, j2 := i, j+1
      if i2 < m && j2 < n {
        result = max(result, abs(board[i][j] - board[i2][j2]))
      }
    }
  }
  return result
}
