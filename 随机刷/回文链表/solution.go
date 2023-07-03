package leetcode

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	fast := head
	slow := head
	stack := []int{}
	for fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			slow = slow.Next
			break
		}
	}
	for i:= len(stack)-1; i>=0; i-- {
		if stack[i] != slow.Val {
			return false
		}
		slow = slow.Next
	}

	return slow == nil
}

