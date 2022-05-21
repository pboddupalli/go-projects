package intervals

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMinMeetinRooms(t *testing.T) {
  intervals := [][]int{{0,30},{5,10},{15,20}}
  assert.Equal(t, 2, minMeetingRooms(intervals))

  intervals = [][]int{{7,10}, {2,4}}
  assert.Equal(t, 1, minMeetingRooms(intervals))
}
