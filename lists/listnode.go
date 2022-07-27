package lists

import "fmt"

type ListNode struct {
  Val int
  Next *ListNode
}

func NewSinglyLL(nums []int) *ListNode {
  dummy := &ListNode{}
  head := dummy
  for _, num := range nums {
    head.Next = &ListNode{Val: num}
    head = head.Next
  }
  return dummy.Next
}

func (head *ListNode) PrintList() {
  for head != nil {
    fmt.Printf("%6d", head.Val)
    head = head.Next
  }
  fmt.Println()
}
