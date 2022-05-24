package arrays

func maxSumSubarray(nums []int) int {
  max, sum := nums[0], nums[0]
  for _, num := range nums[1:] {
    if sum + num < num {
      sum = num
    } else {
      sum += num
    }
    if max < sum {
      max = sum
    }
  }
  return max
}
