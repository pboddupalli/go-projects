package arrays

func subArraySum(nums []int, k int) int {
  hashMap := make(map[int]int)
  hashMap[0] = 1
  cnt,sum := 0, 0
  for _, num := range nums {
    sum += num
    cnt += hashMap[sum-k]
    hashMap[sum]++
  }
  return cnt
}
