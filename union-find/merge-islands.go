package union_find

func numIslands(grid [][]byte) int {
  numIslands := 0
  m,n := len(grid), len(grid[0])
  forest := make([][]int, m)
  for i := 0; i < m; i++ {
    forest[i] = make([]int, n)
  }

  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      if grid[i][j] == '0' {
        continue
      }

      numIslands++
      forest[i][j] = n * i + j // i am my own parent

      if i == 0 && j == 0 {
        // first cell...continue
        continue
      }
      if i > 0 && grid[i-1][j] != '0' {
        merged := merge(forest, i, j, i-1, j)
        if merged {
          numIslands--
        }
      }
      if j > 0 && grid[i][j-1] != '0' {
        merged := merge(forest, i, j, i, j-1)
        if merged {
          numIslands--
        }
      }
    }
  }
  return numIslands
}

func FindParent(forest [][]int, i, j int) int {
  n := len(forest[0])
  px, py := forest[i][j] / n, forest[i][j] % n
  if px != i || py != j {
    forest[i][j] = FindParent(forest, px, py)
  }
  return forest[i][j]
}

func merge(forest [][]int, x1, y1, x2, y2 int) bool {
  n := len(forest[0])
  parent1 := FindParent(forest, x1, y1)
  parent2 := FindParent(forest, x2, y2)
  if parent1 != parent2 {
    forest[parent1 / n][parent1 % n] = parent2
    return true
  }
  return false
}
