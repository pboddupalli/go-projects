package ds

import "fmt"

type Node struct {
  Key int
  Val int
  Prev *Node
  Next *Node
}

type LRUCache struct {
  cap int
  len int
  head *Node
  tail *Node
  hashmap map[int]*Node
}

func Constructor(capacity int) LRUCache {
  res := LRUCache{}
  res.cap = capacity
  res.hashmap = make(map[int]*Node)
  // no dummy for now...so head and tail will be nil
  return res
}

func (this *LRUCache) Get(key int) int {
  node, ok := this.hashmap[key]
  if !ok {
    return -1
  }
  this.moveToFront(node)
  return node.Val
}

func (this *LRUCache) moveToFront(node *Node) {
  if this.head == node {
    // nothing to do
    return
  }
  // since this is not the head, there is a prev element
  node.Prev.Next = node.Next
  if node.Next != nil {
    // non-tail
    node.Next.Prev = node.Prev
  } else {
    // tail node
    this.tail = node.Prev
  }
  node.Prev = nil
  node.Next = this.head
  node.Next.Prev = node
  this.head = node
}

func (this *LRUCache) Put(key int, value int)  {
  // we assume that capacity is at least 1
  node, ok := this.hashmap[key]
  if ok {
    node.Val = value
    this.moveToFront(node)
    return
  }
  if this.len == this.cap {
    // we just replace the tail values and move the tail to the front
    delete(this.hashmap, this.tail.Key)
    this.tail.Key = key
    this.tail.Val = value
    this.hashmap[key] = this.tail
    this.moveToFront(this.tail)
  } else {
    node := &Node{Key: key, Val: value, Next: this.head}
    if this.len == 0 {
      this.tail = node
    } else {
      this.head.Prev = node
    }
    this.head = node
    this.hashmap[key] = node
    this.len++
  }
}

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


