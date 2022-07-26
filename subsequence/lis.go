package subsequence

import "sort"

func lis(nums []int) int {
  result := make([]int, len(nums))
  result[0] = 1
  max := 1
  for i := 1; i < len(nums); i++ {
    result[i] = 1
    for j := i-1; j >= 0; j-- {
      if nums[j] < nums[i] && result[j] + 1 > result[i] {
        result[i] = result[j] + 1
      }
    }
    if result[i] > max {
      max = result[i]
    }
  }
  return max
}

func lisVer2(nums []int) {
  var d []int
  for _, num := range nums {
    if len(d) == 0 || d[len(d)-1] < num {
      d = append(d, num)
    } else {
      p := sort.SearchInts(d, num)
      d[p] = num
    }
  }
}
