package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := new(ListNode)
	low := new(ListNode)
	fast = head
	low = head
	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
			low = low.Next
		} else {
			return low
		}
	}
	return low
}
