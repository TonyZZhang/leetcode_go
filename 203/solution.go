package leetcode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newList := new(ListNode)
	newList = head
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return newList
}
