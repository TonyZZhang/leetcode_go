package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}
//[1,2,3,3,4,4,5]
//[1,1,1,2,3]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head
	jumpFirstFlag := 0
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head = head.Next
			jumpFirstFlag = 1
		} else{
			break
		}
	}

	if jumpFirstFlag == 1 {
		if head.Next == nil {
			return nil
		}
		newHead = head.Next
	}


	for head != nil && head.Next != nil{
		temp := head.Next
		jumpFlag := 0
		if temp.Next == nil {
			break
		}
		for temp.Next.Val == temp.Val {
			temp = temp.Next
			jumpFlag = 1
			if temp == nil || temp.Next == nil {
				break
			}
		}
		if jumpFlag == 1 {
			head.Next = temp.Next
		} else {
			head = head.Next
		}
	}

	return newHead
}