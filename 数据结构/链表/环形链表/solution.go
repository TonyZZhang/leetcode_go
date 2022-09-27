package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	low := new(ListNode)
	fast := new(ListNode)
	low = head
	fast = head
	for low != nil && fast != nil {
		low = low.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		if low == fast {
			return true
		}
	}
	return false
}
