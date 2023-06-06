package leetcode


  type ListNode struct {
      Val int
      Next *ListNode
 }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	newHeadA := headA
	newHeadB := headB
	for newHeadA != nil && newHeadB != nil {
		newHeadA = newHeadA.Next
		newHeadB = newHeadB.Next
	}

	if newHeadA != nil {
		for newHeadA != nil {
			newHeadA = newHeadA.Next
			headA = headA.Next
		}
	}

	if newHeadB != nil {
		for newHeadB != nil {
			newHeadB = newHeadB.Next
			headB = headB.Next
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
