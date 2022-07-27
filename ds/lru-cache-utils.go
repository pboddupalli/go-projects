package ds

import "fmt"

func (this *LRUCache) walkBack() {
  tail := this.tail
  if tail == nil {
    fmt.Println("tail is nil")
  }
  for tail != nil {
    fmt.Printf("...%d", tail.Val)
    tail = tail.Prev
  }
  fmt.Println()
}

func (this *LRUCache) walk() {
  head := this.head
  for head != nil {
    fmt.Printf("...%d", head.Val)
    head = head.Next
  }
  fmt.Println()
}

