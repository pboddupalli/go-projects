package arrays

import (
  "math/rand"
)

type elem struct {
  num, count int
}

func topKFrequent(nums []int, k int) []int {

  m := make(map[int]int)
  for _, num := range nums {
    m[num]++
  }

  counts := make([]elem, 0)
  for k, v := range m {
    counts = append(counts, elem{num: k, count: v})
  }
  // fmt.Println(counts)

  next, i, j := len(counts), 0, len(counts)
  for next != k-1 {
    // fmt.Printf("next: %d i: %d j: %d idx: %d\n", next, i, j, k)

    if next > k-1 {
      j = next
    } else {
      i = next+1
    }

    pivotIdx := i + rand.Intn(j-i)
    pivot := counts[pivotIdx].count
    counts[pivotIdx], counts[j-1] = counts[j-1], counts[pivotIdx]
    next = i

    for idx := i; idx < j; idx++ {
      if counts[idx].count > pivot {
        counts[next], counts[idx] = counts[idx], counts[next]
        next++
      }
    }
    counts[next], counts[j-1] = counts[j-1], counts[next]
  }

  result := make([]int, 0, k)
  for i := 0; i < k; i++ {
    result = append(result, counts[i].num)
  }

  return result
}
