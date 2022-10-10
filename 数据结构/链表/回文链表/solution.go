package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//1-2
//1-2-1
//1-2-2-1
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	newHead := new(ListNode)
	newHead = head
	var old []int
	for head != nil {
		old = append(old, head.Val)
		head = head.Next
	}
	newList := reverseList(newHead)
	for _, v := range old {
		if v != newList.Val {
			return false
		} else {
			newList = newList.Next
		}
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var preHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = preHead
		preHead = head
		head = next
	}
	return preHead
}
