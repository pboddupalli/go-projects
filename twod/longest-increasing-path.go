package twod

var directions [][]int
var maxLength [][]int

func longestIncreasingPath(matrix [][]int) int {
  result := 0
  directions = [][]int{{-1,0}, {1, 0}, {0,1}, {0, -1}}
  maxLength = make([][]int, len(matrix))
  for i := 0; i < len(matrix); i++ {
    maxLength[i] = make([]int, len(matrix[0]))
  }
  m, n := len(matrix), len(matrix[0])

  for x := 0; x < m; x++ {
    for y := 0; y < n; y++ {
      minNeighbor := matrix[x][y]
      for _, dir := range directions {
        x1, y1 := x + dir[0], y + dir[1]
        if x1 < 0 || x1 == m || y1 < 0 || y1 == n {
          continue
        }
        minNeighbor = min(minNeighbor, matrix[x1][y1])
      }
      if minNeighbor < matrix[x][y] {
        // that guy will take care of determining.
        continue
      }
      result = max(result, dfs(matrix, x, y))
    }
  }
  return result
}

func dfs(nums [][]int, x, y int) int {
  if maxLength[x][y] != 0 {
    return maxLength[x][y]
  }
  result := 1
  for _, dir := range directions {
    x1, y1 := x + dir[0], y + dir[1]
    if x1 < 0 || x1 == len(nums) || y1 < 0 || y1 == len(nums[0]) {
      continue
    }
    if nums[x1][y1] <= nums[x][y] {
      continue
    }
    result = max(result, 1 + dfs(nums, x1, y1))
  }
  maxLength[x][y] = result
  return result
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
