package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//找入口的推断
//https://zhuanlan.zhihu.com/p/410356891
//相遇之后，快指针重新回到起点，然后快指针步长改为1，两者再次相遇的时候，就是入口
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
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
			return nil
		}
		if low == fast {
			fast = head
			for fast != low {
				low = low.Next
				fast = fast.Next
			}
			return low
		}
	}
	return nil
}
