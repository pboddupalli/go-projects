package arrays

import (
  "math"
  "sort"
)

func findMinDifference(timePoints []string) int {
  result := math.MaxInt
  sort.Strings(timePoints)

  for i := 1; i < len(timePoints); i++ {
    prev := toMinutes(timePoints[i-1])
    current := toMinutes(timePoints[i])
    if current - prev < result {
      result = current - prev
    }
  }

  prev := toMinutes(timePoints[len(timePoints)-1])
  current := toMinutes(timePoints[0]) + 24 * 60
  if current - prev < result {
    result = current - prev
  }

  return result
}

func toMinutes(t string) int {
  h := (t[0] - '0') * 10 + (t[1] - '0')
  m := (t[3] - '0') * 10 + (t[4] - '0')
  return int(h) * 60 + int(m)
}
