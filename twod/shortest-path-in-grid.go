package twod

type state struct {
  x, y int
  k int
}

func shortestPath(grid [][]int, k int) int {
  m, n := len(grid), len(grid[0])
  if m == 1 && n == 1 {
    return 0
  }
  directions := [][]int { {0, 1}, {0,-1}, {1, 0}, {-1, 0} }

  leftK := make([][]int, m)
  for i := 0; i < m; i++ {
    leftK[i] = make([]int, n)
    for j := 0; j < n; j++ {
      leftK[i][j] = -1
    }
  }

  Q := make([]state, 0)
  Q = append(Q, state{x: 0, y: 0, k: k})
  steps := 0

  for len(Q) > 0 {

    nextSet := Q
    Q = Q[len(Q):] // empty it out
    steps++

    for _, elem := range nextSet {
      credits := elem.k
      if grid[elem.x][elem.y] == 1 {
        credits--
      }
      if credits < 0 {
        continue
      }
      for _, dir := range directions {
        x1, y1 := elem.x + dir[0], elem.y + dir[1]
        if x1 < 0 || x1 == m || y1 < 0 || y1 == n {
          continue
        }
        if x1 == m-1 && y1 == n-1 {
          return steps
        }
        if leftK[x1][y1] == -1 || leftK[x1][y1] < credits {
          leftK[x1][y1] = credits
          Q = append(Q, state{x: x1, y:y1, k: credits})
        }
      }
    }
  }
  return -1
}
