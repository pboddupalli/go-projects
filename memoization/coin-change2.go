package memoization

func change(amount int, coins []int) int {
  dp := make([][]int, len(coins))
  for i := 0; i < len(coins); i++ {
    dp[i] = make([]int, amount + 1)
  }
  retval := getChange(amount, coins, 0, dp)
  if retval == -1 {
    return 0
  }
  return retval
}

func getChange(amount int, coins []int, idx int, dp [][]int) int {
  if amount == 0 {
    return 1
  }
  if idx == len(coins) {
    return 0
  }
  if dp[idx][amount] == -1 {
    return 0
  } else if dp[idx][amount] > 0 {
    return dp[idx][amount]
  }
  result := 0
  for i := 0; amount - i * coins[idx] >= 0;  i++ {
    tmp := getChange(amount - i * coins[idx], coins, idx + 1, dp)
    if tmp != -1 {
      result += tmp
    }
  }
  if result == 0 {
    result = -1
  }
  dp[idx][amount] = result
  return result
}
