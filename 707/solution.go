package leetcode

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val int
	Next *Node
}

func Constructor() MyLinkedList {
	dummyHead := new(Node)
	dummyHead.Val = -1
	dummyHead.Next = nil

	return MyLinkedList{
		dummy: dummyHead,
	}
}


func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	for i:= 0; i< index; i++ {
		if head == nil {
			return -1
		}
		head = head.Next
	}
	return head.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	newHead := new(Node)
	newHead.Val = val
	newHead.Next = this.dummy.Next
	this.dummy.Next = newHead
}


func (this *MyLinkedList) AddAtTail(val int)  {
	head := this.dummy.Next
	last := new(Node)
	tailNode := new(Node)
	tailNode.Val = val
	tailNode.Next = nil

	for head != nil {
		last = head
		head = head.Next
	}
	last.Next = tailNode
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.GetListLength() {
		return
	}

	if index == this.GetListLength() {
		this.AddAtTail(val)
	}

	if index <=0 {
		this.AddAtHead(val)
	}

	insertNode := new(Node)
	insertNode.Val = val
	insertNode.Next = nil

	head := this.dummy.Next
	for i := 0; i < index; i++ {
		if head == nil ||head.Next == nil {
			break
		}
		head = head.Next
	}
	insertNode.Next = head.Next
	insertNode = head.Next
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index >= this.GetListLength() || index <0 {
		return
	}

	if index == 0 {
		this.dummy.Next = this.dummy.Next.Next
		return
	}

	preNode := new(Node)
	head := this.dummy.Next
	for i:=0; i < index; i++ {
		if head == nil || head.Next == nil {
			break
		}
		preNode = head
		head = head.Next
	}

	preNode.Next = head.Next
}

func (this *MyLinkedList) GetListLength() int {
	head := this.dummy.Next
	length := 0

	for head != nil {
		length++
		head = head.Next
	}

	return length
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
