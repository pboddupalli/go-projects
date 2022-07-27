package ds

type Node struct {
  Key   int
  Value int
  Next *Node
}

type LRUCache struct {
  Dummy           *Node
  Tail            *Node
  Capacity        int
  KeyToPrevMap   map[int]*Node
}


func Constructor(capacity int) LRUCache {
  keyToPrevMap := make(map[int]*Node)
  dummy := &Node{}
  lruCache := LRUCache{
    Dummy: dummy,
    Tail: dummy,
    Capacity: capacity,
    KeyToPrevMap: keyToPrevMap,
  }

  return lruCache
}


func (this *LRUCache) Get(key int) int {
  // fmt.Println("Get....")
  prevNode, ok := this.KeyToPrevMap[key]
  if !ok {
    return -1
  }

  node := prevNode.Next
  value := node.Value
  // if node != this.Tail {
  this.Kick(prevNode)
  // this.PushBack(node)
  // }

  return value
}


func (this *LRUCache) Put(key int, value int)  {
  prevNode, ok := this.KeyToPrevMap[key]
  if ok {
    node := prevNode.Next
    node.Value = value
    // if node != this.Tail {
    this.Kick(prevNode)
    // this.PushBack(node)
    // }
    return
  }
  newNode := &Node{
    Key: key,
    Value: value,
  }
  this.PushBack(newNode)

  if len(this.KeyToPrevMap) > this.Capacity {
    this.PopFront()
  }
}

func (this *LRUCache) Kick(prevNode *Node) {
  node := prevNode.Next
  if node == this.Tail {
    return
  }

  prevNode.Next = node.Next
  this.KeyToPrevMap[node.Next.Key] = prevNode
  this.PushBack(node)
}

func (this *LRUCache) PushBack(node *Node) {
  this.Tail.Next = node
  node.Next = nil
  this.KeyToPrevMap[node.Key] = this.Tail
  this.Tail = node
}

func (this *LRUCache) PopFront() {
  node := this.Dummy.Next
  this.Dummy.Next = node.Next
  delete(this.KeyToPrevMap, node.Key)
  this.KeyToPrevMap[node.Next.Key] = this.Dummy
}
