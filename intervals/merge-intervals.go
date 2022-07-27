package intervals

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
  var result [][]int
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })
  result = append(result, intervals[0])
  for _, next := range intervals[1:] {
    if next[0] > result[len(result)-1][1] {
      result = append(result, next)
    } else if next[1] > result[len(result)-1][1] {
      result[len(result)-1][1] = next[1]
    }
  }
  return result
}

func mergeIntervalsVer2(intervals [][]int) [][]int {
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })
  current := 0
  for i := 1; i < len(intervals); i++ {
    if intervals[i][0] > intervals[current][1] {
      current++
      intervals[current][0], intervals[current][1] = intervals[i][0], intervals[i][1]
    } else {
      intervals[current][1] = max(intervals[current][1], intervals[i][1])
    }
  }
  return intervals[:current+1]
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}