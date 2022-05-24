package memoization

import (
  "fmt"
  "math"
  "sort"
)

func coinChange2(coins []int, amount int) int {
  if amount == 0 {
    return 0
  }
  count := make(map[int]int)
  minCoin := math.MaxInt
  for _, coin := range coins {
    if coin < minCoin {
      minCoin = coin
    }
    count[coin] = 1
  }
  return getCoinChange2(amount, count, minCoin)
}

func getCoinChange2(amount int, count map[int]int, minCoin int) int {
  result := -1
  val, ok := count[amount]
  if ok {
    return val
  }
  for i := minCoin; i <= amount/2; i++ {
    result1 := getCoinChange2(i, count, minCoin)
    result2 := getCoinChange2(amount-i, count, minCoin)
    if result1 != -1 && result2 != -1 {
      if result == -1 || result > result1 + result2 {
        result = result1 + result2
      }
    }
  }
  if result != -1 {
    count[amount] = result
  }
  return result
}

func coinChange(coins []int, amount int) int {
  sort.Slice(coins, func(i, j int) bool { return coins[i] > coins[j]} )
  fmt.Println(coins)
  memo := make([]int, amount+1)
  return getCoinChange(amount, coins, 0, memo)
}

func getCoinChange(amount int, coins []int, idx int, memo []int) int {
  if amount == 0 {
    return 0
  }
  if idx == len(coins) {
    return -1
  }
  if memo[amount] != 0 {
    return memo[amount]
  }
  result := -1
  for i := 0; amount - i * coins[idx] >= 0; i++ {
    tmp := getCoinChange(amount - i * coins[idx], coins, idx + 1, memo)
    if tmp != -1 && (result == -1 || result > i + tmp) {
      result = i + tmp
    }
  }
  memo[amount] = result
  return result
}
