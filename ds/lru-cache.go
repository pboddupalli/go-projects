package ds

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

  // Step 1: Removal from current Position
  this.removeFromList(node)

  // Step 2: Add at the head of the list
  this.insertAtHead(node)
}

func (this *LRUCache) Put(key int, value int)  {
  // we assume that capacity is at least 1
  node, ok := this.hashmap[key]
  if ok {
    node.Val = value
    this.moveToFront(node)
  }
  if !ok {
    if this.len == this.cap {
      delete(this.hashmap, this.tail.Key)
      node = this.tail
      node.Key = key
      node.Val = value
      this.moveToFront(node)
      this.hashmap[key] = node
    } else {
      node = &Node{Key:key, Val: value}
      if this.len == 0 {
        this.tail = node
      }
      this.insertAtHead(node)
      this.hashmap[key] = node
      this.len++
    }
  }
}

func (this *LRUCache) removeFromList(node *Node) {
  if node.Prev != nil {
    node.Prev.Next = node.Next
  }
  if node.Next != nil {
    node.Next.Prev = node.Prev
  }
  if node == this.tail {
    this.tail = node.Prev
  }
}

func (this *LRUCache) insertAtHead(node *Node) {
  node.Prev = nil
  node.Next = this.head
  if this.head != nil {
    this.head.Prev = node
  }
  this.head = node
}
