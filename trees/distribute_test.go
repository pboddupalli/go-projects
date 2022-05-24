package trees

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestDistribute(t *testing.T) {
  root := BuildBinaryTree([]interface{}{0,3,0})
  assert.Equal(t, 3, distributeCoins(root))

  root = BuildBinaryTree([]interface{}{3,0,0})
  assert.Equal(t, 2, distributeCoins(root))
}
