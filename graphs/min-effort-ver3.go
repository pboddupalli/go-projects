package graphs

import "math"

func minimumEffortPathVer3(board [][]int) int {

  m,n := len(board), len(board[0])

  seen := make([][]int, m)
  for i := 0; i < m; i++ {
    seen[i] = make([]int, n)
    for j := 0; j < n; j++ {
      seen[i][j] = -1
    }
  }
  return dfsVer3(board, seen, 0, 0)
}

func dfsVer3(board [][]int, seen [][]int, x, y int) int {
  directions := [][]int{{-1,0}, {1,0}, {0,1},{0,-1}}
  m, n := len(seen), len(seen[0])
  if x == m-1 && y == n-1 {
    return 0
  }
  if seen[x][y] != -1 {
    return seen[x][y]
  }
  seen[x][y] = math.MaxInt
  for _, dir := range directions {
    x1, y1 := x + dir[0], y + dir[1]
    if x1 < 0 || x1 >= m || y1 < 0 || y1 >= n {
      continue
    }
    if seen[x1][y1] == math.MaxInt {
      // dont go into child that is already visited in this run
      continue
    }
    childResult := dfsVer3(board, seen, x1, y1)
    edgeValue := abs(board[x][y] - board[x1][y1])
    seen[x][y] = min(seen[x][y], max(edgeValue, childResult))
  }
  result := seen[x][y]
  seen[x][y] = -1
  return result
}
