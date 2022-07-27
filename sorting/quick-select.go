package sorting

import (
  "math/rand"
)

// returns the pivot index
// always cover corner cases
func quickSelect(nums []int) int {
  if len(nums) == 0 {
    return -1
  }
  seed := rand.NewSource(43)
  rd := rand.New(seed)
  idx := rd.Intn(len(nums))
  nums[0], nums[idx] = nums[idx], nums[0]
  i, j := 1, len(nums)-1
  for i <= j {
    for i < len(nums) && nums[i] < nums[0] { i++ }
    for j >= 0 && nums[j] >= nums[0] { j-- }
    if i < j {
      nums[i], nums[j] = nums[j], nums[i]
      i++
      j--
    }
  }
  nums[i-1], nums[0] = nums[0], nums[i-1]
  return i-1
}

func quickSelect2(nums []int) int {
  if len(nums) == 0 {
    return -1
  }
  next := 0
  for i := 0; i < len(nums); i++ {
    if nums[i] > nums[len(nums)-1] {
      nums[i], nums[next] = nums[next], nums[i]
      next++
    }
  }
  nums[next], nums[len(nums)-1] = nums[len(nums)-1], nums[next]
  return next
}
