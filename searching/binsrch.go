package searching

func binSrch(nums []int, target int) int {
  i, j := 0, len(nums)
  for i < j {
    mid := i + (j - i)/2
    if nums[mid] == target {
      return mid
    } else if target > nums[mid] {
      i = mid + 1
    } else {
      j = mid
    }
  }
  return -1
}
