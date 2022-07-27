package containers

func maxArea(height []int) int {
  i, j := 0, len(height)-1
  result := 0
  for i < j {
    current := min(height[i], height[j]) * (j - i)
    if current > result {
      result = current
    }
    if height[i] > height[j] {
      j--
    } else {
      i++
    }
  }
  return result
}

func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}
