package hard

import (
  "fmt"
  "sort"
)

func russianDolls(dolls [][]int) int {
  sort.Slice(dolls, func (i, j int) bool {
    if dolls[i][0] == dolls[j][0] { // if heights are equal, compare weights
      return dolls[i][1] > dolls[j][1]
    }
    return dolls[i][0] < dolls[j][0]
  })
  fmt.Println(dolls)
  dp := make([]int, 0)
  dp = append(dp, dolls[0][1])
  for i := 1; i < len(dolls); i++ {
    if dolls[i][1] > dp[len(dp)-1] {
      dp = append(dp, dolls[i][1])
    } else {
      idx := binSrch(dp, dolls[i][1])
      dp[idx] = dolls[i][1]
    }
  }
  return len(dp)
}

// returns the index of the first element that is either equal
// to or greater than target
func binSrch(nums []int, target int) int {
  i, j := 0, len(nums)
  result := -1
  for i < j {
    mid := i + (j - i)/2
    if target == nums[mid] {
      return mid
    }
    if nums[mid] > target { // potential one
      result = mid
      j = mid
    } else {
      i = mid + 1
    }
  }
  return result
}
