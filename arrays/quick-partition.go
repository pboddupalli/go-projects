package arrays

import "math/rand"

func quickPartition(nums []int) int {
  pivotIdx := rand.Intn(len(nums)) // the value will be between 0 and len(nums)-1
  pivot := nums[pivotIdx]
  nums[len(nums)-1], nums[pivotIdx] = nums[pivotIdx], nums[len(nums)-1]
  next := 0
  for i := 0; i < len(nums); i++ {
    if nums[i] > pivot {
      nums[i], nums[next] = nums[next], nums[i]
      next++
    }
  }
  nums[next], nums[len(nums)-1] = nums[len(nums)-1], nums[next]
  return next
}
