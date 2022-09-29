package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	newHeadA := new(ListNode)
	newHeadB := new(ListNode)
	newHeadA = headA
	newHeadB = headB
	for newHeadA != nil && newHeadB != nil {
		if newHeadA == newHeadB {
			return newHeadA
		}
		newHeadA = newHeadA.Next
		newHeadB = newHeadB.Next
		if newHeadA == nil && newHeadB != nil {
			newHeadA = headB
		}
		if newHeadB == nil && newHeadA != nil {
			newHeadB = headA
		}
	}
	return nil
}
