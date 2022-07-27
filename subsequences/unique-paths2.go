package subsequences

func uniquePaths2(grid [][]int) int {
  m, n := len(grid), len(grid[0])
  prev := make([]int, n+1)
  current := make([]int, n+1)
  current[1] = 1 // because grid[0][0] is never zero
  for i := 0; i < m; i++ {
    prev, current = current, prev
    for j := 1; j <= n; j++ {
      if grid[i][j-1] == 1 {
        current[j] = 0
      } else {
        current[j] = current[j-1] + prev[j]
      }
    }
  }
  return current[n]
}

