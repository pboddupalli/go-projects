package caching

import "fmt"

type Node struct {
  key int
  val int
  prev *Node
  next *Node
}

type LRUCache struct {
  cap int
  len int
  head *Node
  tail *Node
  dict map[int]*Node
}

func Constructor(capacity int) LRUCache {
  result := LRUCache{}
  result.cap = capacity
  result.dict = make(map[int]*Node)
  dummy := &Node{}
  dummy.prev, dummy.next = dummy, dummy
  result.head, result.tail = dummy, dummy
  return result
}

func (this *LRUCache) Get(key int) int {
  node, ok := this.dict[key]
  if !ok {
    // nothing to do
    return -1
  }
  this.moveToFront(node)
  return node.val
}

func (this *LRUCache) moveToFront(node *Node) {
  if this.head == node {
    return
  }
  node.prev.next = node.next
  node.next.prev = node.prev
  node.next = this.head
  this.head = node
}

func (this *LRUCache) Put(key int, value int)  {
  node, ok := this.dict[key]
  if ok {
    node.val = value
    this.moveToFront(node)
    return
  }
  if this.len < this.cap {
    node = &Node{next: this.head}
    this.head.prev = node
    this.head = node
    this.len++
  } else {
    delete(this.dict, this.tail.prev.key)
    node = this.tail.prev
    this.moveToFront(node)
  }
  node.key = key
  node.val = value
  this.dict[key] = node
}

func (this *LRUCache) info() {
  fmt.Println(this.len)
  fmt.Println(this.cap)
  if this.head != nil {
    fmt.Println("head: ", this.head.val)
  }
  if this.tail != nil {
    fmt.Println("tail: ", this.tail.val)
  } else {
    fmt.Println("tail is nil")
  }
  head := this.head
  for head != nil {
    fmt.Printf("key: %2d value: %2d\t", head.key, head.val)
    head = head.next
  }
  fmt.Println()
}
