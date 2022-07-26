package arrays

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
  hashMap := make(map[int]int)
  hashMap[nums[0] % k] = 0
  sum := nums[0]
  for i := 1; i < len(nums); i++ {
    sum += nums[i]
    mod := sum % k
    if mod == 0 {
      return true
    }
    fmt.Println(sum, mod, hashMap)
    idx, ok := hashMap[mod]
    if ok {
      if (i - idx) > 1 {
        return true
      }
    } else {
      hashMap[mod] = i
    }
  }
  return false
}
