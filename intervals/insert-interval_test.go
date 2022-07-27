package intervals

import (
  "fmt"
  "testing"
)

func TestInsertInterval(t *testing.T) {
  intervals := [][]int{{1,3},{6,9}}

  fmt.Println(insert(intervals, []int{2,5}))

  intervals = [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
  fmt.Println(insert(intervals, []int{4,8}))
}
