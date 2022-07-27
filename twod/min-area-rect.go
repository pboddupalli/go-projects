package twod

import (
  "math"
  "sort"
)

type ymap struct {
  y int
  xcoords []int
}

func minAreaRect(points [][]int) int {
  m := make(map[int][]int)
  for _, point := range points {
    m[point[1]] = append(m[point[1]], point[0])
  }
  for _, v := range m {
    sort.Ints(v)
  }
  ymaps := make([]ymap, 0)
  for k, v := range m {
    ymaps = append(ymaps, ymap{k, v})
  }
  sort.Slice(ymaps, func(i, j int) bool { return ymaps[i].y > ymaps[j].y})

  result := math.MaxInt
  for i := 0; i < len(ymaps)-1; i++ {
    for j := 0; j < len(ymaps[i].xcoords)-1; j++ {
      for k := j+1; k < len(ymaps[i].xcoords); k++ {
        x1, x2 := ymaps[i].xcoords[j], ymaps[i].xcoords[k]
        for l := i+1; l < len(ymaps); l++ {
          if binSrch(ymaps[l].xcoords, x1) && binSrch(ymaps[l].xcoords, x2) {
            result = min(result, (ymaps[i].y - ymaps[l].y) * (x2 - x1))
          }
        }
      }
    }
  }

  return result
}

func binSrch(nums []int, target int) bool {
  i, j := 0, len(nums)
  for i < j {
    mid := i + (j - i) / 2
    if nums[mid] == target {
      return true
    }
    if target < nums[mid] {
      j = mid
    } else {
      i = mid + 1
    }
  }
  return false
}
