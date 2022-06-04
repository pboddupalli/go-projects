package lists

func mergeKLists(lists []*ListNode) *ListNode {
  if len(lists) == 0 {
    return nil
  }
  for len(lists) > 1 {
    next := 0
    for j := 0; j+1 < len(lists); j = j + 2 {
      retval := merge(lists[j], lists[j+1])
      lists[next] = retval
      next++
    }
    if len(lists) % 2 == 1 {
      lists[next] = lists[len(lists)-1]
      next++
    }
    lists = lists[:next]
  }
  return lists[0]
}

func merge(list1, list2 *ListNode) *ListNode {
  result := &ListNode{}
  current := result
  for list1 != nil && list2 != nil  {
    if list1.Val <= list2.Val {
      current.Next = list1
      list1 = list1.Next
    } else {
      current.Next = list2
      list2 = list2.Next
    }
    current = current.Next
  }
  if list1 != nil {
    current.Next = list1
  } else {
    current.Next = list2
  }
  return result.Next
}
