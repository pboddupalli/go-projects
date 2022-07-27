package arrays

import (
  "fmt"
  "math/rand"
  "testing"
)

func TestQuickPartition(t *testing.T) {
  nums := make([]int, 0)
  for i := 0; i < 16; i++ {
    nums = append(nums, rand.Intn(64))
  }
  fmt.Println(nums)
  pivot := quickPartition(nums)
  fmt.Println(pivot, nums)
}
