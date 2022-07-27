package toposort

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCourseScheduleOne(t *testing.T) {
  assert.Equal(t, false, canFinish(2, [][]int{{0,1}, {1,0}}))
}
