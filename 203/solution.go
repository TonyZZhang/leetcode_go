package leetcode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	all := false
	newHead := new(ListNode)
	for head != nil {
		if head.Val != val {
			newHead = head
			all = true
			break
		} else {
			head = head.Next
		}
	}

	if !all {
		return nil
	}

	for head != nil {
		if head.Next != nil {
			if head.Next.Val == val {
				head.Next = head.Next.Next
			} else {
				head = head.Next
			}
		}
	}
	return newHead
}
