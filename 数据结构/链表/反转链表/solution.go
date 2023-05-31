package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//preHead := new(ListNode)
	var preHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = preHead
		preHead = head
		head = next
	}
	return preHead
}

func reverseListV2(l *ListNode) {
	empty := &ListNode{}
	for l != nil && l.Next != nil {
		cur := l
		next := l.Next
		l.Next = empty
		l = next
		empty = cur
	}
}