package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	mid := findMidNode(head)
	left := sortList(head)
	right := sortList(mid)

	return mergeTwoList(left, right)
}

func findMidNode(head *ListNode) *ListNode {
	//找中点
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {

}
