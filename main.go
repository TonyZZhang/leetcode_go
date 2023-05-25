package main

import (
	"fmt"
)

func main(){
	head := new(ListNode)
	head.Val = 1
	head.Next = nil
	head.addNode(2)
	head.addNode(3)
	for head!= nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) addNode(value int) {
	newHead := new(ListNode)
	newHead = l
	last  := new(ListNode)
	last.Val = value
	last.Next = nil
	if newHead != nil {
		for newHead != nil {
			if newHead.Next != nil {
				newHead = newHead.Next
			}else {
				newHead.Next = last
				break
			}
		}
	}
}
