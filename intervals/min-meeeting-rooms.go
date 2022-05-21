package intervals

import "sort"

func minMeetingRooms(intervals [][]int) int {
  start := make([]int, len(intervals))
  end := make([]int, len(intervals))
  for i, interval := range intervals {
    start[i] = interval[0]
    end[i] = interval[1]
  }
  sort.Ints(start)
  sort.Ints(end)
  numRooms, result := 0,0
  for i, j := 0,0; i < len(start); {
    if start[i] < end[j] {
      numRooms++
      if numRooms > result {
        result = numRooms
      }
      i++
    } else if end[j] < start[i] {
      numRooms--
      j++
    } else {
      i++; j++
    }
  }
  return result
}
