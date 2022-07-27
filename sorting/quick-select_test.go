package sorting

import (
  "fmt"
  "math/rand"
  "testing"
)

func TestQuickSelect(t *testing.T) {
  nums := []int{}
  for i := 0; i < 10; i++ {
    nums = append(nums, rand.Intn(128))
  }
  nums2 := make([]int, len(nums))
  copy(nums2, nums)

  result := quickSelect(nums)
  fmt.Println(nums, nums[result])

  result = quickSelect2(nums2)
  fmt.Println(nums2, nums[result])
}
