package leetcode


Definition for singly-linked list.
type ListNode struct {
  Val int
  Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  if n < 1 {
    return head
  }
  newHead := head
  var lens int
  for newHead != nil {
    lens++
    newHead = newHead.Next
  }

  pre := lens-n-1
  if pre < 0 {
    return head.Next
  }
  temp := &ListNode{}
  newHead = head
  for i := 0; i <= lens-n; i++ {
    if i == pre {
      temp = newHead
      break
    }
    newHead = newHead.Next
  }
  temp.Next = temp.Next.Next
  return head
}
