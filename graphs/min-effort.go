package graphs

import (
  "fmt"
  "math"
)

func minimumEffortPathVer2(heights [][]int) int {
  seen := make([][]int, len(heights))
  for i := 0 ;i < len(heights); i++ {
    seen[i] = make([]int, len(heights[0]))
    for j := 0; j < len(heights[0]); j++ {
      seen[i][j] = -1
    }
  }
  retval := dfs(heights, 0,0, seen)
  return retval
}

func dfs(heights [][]int, x, y int, seen [][]int) int {
  m, n := len(heights), len(heights[0])
  directions := [][]int{{-1,0}, {1,0}, {0,1}, {0,-1}}

  if x == m-1 && y == n-1 {
    return 0
  }
  if seen[x][y] >= 0 {
    return seen[x][y]
  }
  seen[x][y] = math.MaxInt
  for _, dir := range directions {
    x1, y1 := x + dir[0], y + dir[1]
    if x == 0 && y == 0 {
      fmt.Println("x1,y1", x1, y1)
    }
    if x1 < 0 || x1 >= m || y1 < 0 || y1 >= n {
      continue
    }
    if seen[x1][y1] == math.MaxInt {
      continue
    }
    retval := dfs(heights, x1, y1, seen)
    seen[x][y] = min(seen[x][y], max(retval, abs(heights[x][y] - heights[x1][y1])))
    if x == 0 && y == 0 {
      fmt.Println("0,0", seen[0][0])
    }
  }
  return seen[x][y]
}

func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}
