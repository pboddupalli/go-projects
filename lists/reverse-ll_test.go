package lists

import (
  "testing"
)

func TestReverseLl(t *testing.T) {
  head := NewSinglyLL([]int{1,2,3})
  head.PrintList()
  head = reverseList(head)
  head.PrintList()
}
